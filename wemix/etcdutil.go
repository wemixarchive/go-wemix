/* etcdutil.go */

package wemix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
	"go.etcd.io/etcd/server/v3/etcdserver/api/membership"
	"go.etcd.io/etcd/server/v3/etcdserver/api/v3client"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	wemixapi "github.com/ethereum/go-ethereum/wemix/api"
)

var (
	etcdLock = &SpinLock{0}
)

func (ma *wemixAdmin) etcdMemberExists(name, cluster string) (bool, error) {
	var node *wemixNode
	ma.lock.Lock()
	for _, i := range ma.nodes {
		if i.Name == name || i.Id == name || i.Ip == name {
			node = i
			break
		}
	}
	ma.lock.Unlock()

	if node == nil {
		return false, ethereum.NotFound
	}
	host := fmt.Sprintf("%s:%d", node.Ip, node.Port+1)

	var ss []string
	if ss = strings.Split(cluster, ","); len(ss) <= 0 {
		return false, ethereum.NotFound
	}

	for _, i := range ss {
		if j := strings.Split(i, "="); len(j) == 2 {
			u, err := url.Parse(j[1])
			if err == nil && u.Host == host {
				return true, nil
			}
		}
	}

	return false, nil
}

// fill the missing name in cluster string when a member is just added, like
// "=http://1.1.1.1:8590,wemix2=http:/1.1.1.2:8590"
func (ma *wemixAdmin) etcdFixCluster(cluster string) (string, error) {
	if ma.self == nil {
		return "", ethereum.NotFound
	}

	host := fmt.Sprintf("%s:%d", ma.self.Ip, ma.self.Port+1)

	var ss []string
	if ss = strings.Split(cluster, ","); len(ss) <= 0 {
		return "", ethereum.NotFound
	}

	var bb bytes.Buffer
	for _, i := range ss {
		if j := strings.Split(i, "="); len(j) == 2 {
			if bb.Len() > 0 {
				bb.WriteString(",")
			}

			if len(j[0]) != 0 {
				bb.WriteString(i)
			} else {
				u, err := url.Parse(j[1])
				if err != nil || u.Host != host {
					bb.WriteString(i)
				} else {
					bb.WriteString(fmt.Sprintf("%s=%s", ma.self.Name, j[1]))
				}
			}
		}
	}

	return bb.String(), nil
}

func (ma *wemixAdmin) etcdNewConfig(newCluster bool) *embed.Config {
	// LPUrls: listening peer urls
	// APUrls: advertised peer urls
	// LCUrls: listening client urls
	// LPUrls: advertised client urls
	cfg := embed.NewConfig()
	cfg.PeerAutoTLS = true
	cfg.ClientAutoTLS = true
	cfg.SelfSignedCertValidity = 10
	cfg.AutoCompactionMode = "revision"
	cfg.AutoCompactionRetention = "100"
	cfg.LogLevel = "error"
	cfg.Dir = ma.etcdDir
	cfg.Name = ma.self.Name
	u, _ := url.Parse(fmt.Sprintf("http://%s:%d", "0.0.0.0", ma.self.Port+1))
	cfg.LPUrls = []url.URL{*u}
	u, _ = url.Parse(fmt.Sprintf("http://%s:%d", ma.self.Ip, ma.self.Port+1))
	cfg.APUrls = []url.URL{*u}
	u, _ = url.Parse(fmt.Sprintf("http://localhost:%d", ma.self.Port+2))
	cfg.LCUrls = []url.URL{*u}
	cfg.ACUrls = []url.URL{*u}
	if newCluster {
		cfg.ClusterState = embed.ClusterStateFlagNew
		cfg.ForceNewCluster = true
	} else {
		cfg.ClusterState = embed.ClusterStateFlagExisting
	}
	cfg.InitialCluster = fmt.Sprintf("%s=http://%s:%d", ma.self.Name,
		ma.self.Ip, ma.self.Port+1)
	cfg.InitialClusterToken = etcdClusterName
	return cfg
}

func (ma *wemixAdmin) etcdIsRunning() bool {
	return ma.etcd != nil && ma.etcdCli != nil
}

func (ma *wemixAdmin) etcdGetCluster() string {
	if !ma.etcdIsRunning() {
		return ""
	}

	var ms []*membership.Member
	ms = append(ms, ma.etcd.Server.Cluster().Members()...)
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Attributes.Name < ms[j].Attributes.Name
	})

	var bb bytes.Buffer
	for _, i := range ms {
		if bb.Len() > 0 {
			bb.WriteString(",")
		}
		bb.WriteString(fmt.Sprintf("%s=%s", i.Attributes.Name,
			i.RaftAttributes.PeerURLs[0]))
	}
	return bb.String()
}

// returns new cluster string if adding the member is successful
func (ma *wemixAdmin) etcdAddMember(name string) (string, error) {
	if !ma.etcdIsRunning() {
		return "", ErrNotRunning
	}

	if ok, _ := ma.etcdMemberExists(name, ma.etcdGetCluster()); ok {
		return ma.etcdGetCluster(), nil
	}

	var node *wemixNode
	ma.lock.Lock()
	for _, i := range ma.nodes {
		if i.Name == name || i.Enode == name || i.Id == name || i.Ip == name {
			node = i
			break
		}
	}
	ma.lock.Unlock()

	if node == nil {
		return "", ethereum.NotFound
	}

	_, err := ma.etcdCli.MemberAdd(context.Background(),
		[]string{fmt.Sprintf("http://%s:%d", node.Ip, node.Port+1)})
	if err != nil {
		log.Error("failed to add a new member",
			"name", name, "ip", node.Ip, "port", node.Port+1, "error", err)
		return "", err
	} else {
		log.Info("a new member added",
			"name", name, "ip", node.Ip, "port", node.Port+1, "error", err)
		return ma.etcdGetCluster(), nil
	}
}

// returns new cluster string if removing the member is successful
func (ma *wemixAdmin) etcdRemoveMember(name string) (string, error) {
	if !ma.etcdIsRunning() {
		return "", ErrNotRunning
	}

	var id uint64
	for _, i := range ma.etcd.Server.Cluster().Members() {
		if i.Attributes.Name == name {
			id = uint64(i.ID)
			break
		}
	}
	if id == 0 {
		id, _ = strconv.ParseUint(name, 16, 64)
		if id == 0 {
			return "", ethereum.NotFound
		}
	}

	_, err := ma.etcdCli.MemberRemove(context.Background(), id)
	if err != nil {
		return "", err
	}

	return ma.etcdGetCluster(), nil
}

func (ma *wemixAdmin) etcdMoveLeader(name string) error {
	if !ma.etcdIsRunning() {
		return ErrNotRunning
	}

	var id uint64
	for _, i := range ma.etcd.Server.Cluster().Members() {
		if i.Attributes.Name == name {
			id = uint64(i.ID)
			break
		}
	}
	if id == 0 {
		id, _ = strconv.ParseUint(name, 16, 64)
		if id == 0 {
			return ethereum.NotFound
		}
	}
	to := 1500 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), to)
	err := ma.etcd.Server.MoveLeader(ctx, ma.etcd.Server.Lead(), id)
	cancel()
	return err
}

func (ma *wemixAdmin) etcdTransferLeadership() error {
	if !ma.etcdIsRunning() {
		return ErrNotRunning
	}
	return ma.etcd.Server.TransferLeadership()
}

func (ma *wemixAdmin) etcdWipe() error {
	if ma.etcdIsRunning() {
		ma.etcdCli.Close()
		ma.etcd.Server.Stop()
		ma.etcd = nil
		ma.etcdCli = nil
	}

	if _, err := os.Stat(ma.etcdDir); err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	} else {
		return os.RemoveAll(ma.etcdDir)
	}
}

func (ma *wemixAdmin) etcdInit() error {
	if ma.etcdIsRunning() {
		return ErrAlreadyRunning
	} else if ma.self == nil {
		return ErrNotRunning
	}

	cfg := ma.etcdNewConfig(true)
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Error("failed to initialize etcd", "error", err)
		return err
	} else {
		log.Info("initialized etcd server")
	}

	ma.etcd = etcd
	ma.etcdCli = v3client.New(etcd.Server)
	return nil
}

func (ma *wemixAdmin) etcdStart() error {
	if ma.etcdIsRunning() {
		return ErrAlreadyRunning
	}

	cfg := ma.etcdNewConfig(false)
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Error("failed to start etcd", "error", err)
		return err
	} else {
		log.Info("started etcd server")
	}

	ma.etcd = etcd
	ma.etcdCli = v3client.New(etcd.Server)

	// capture leader changes & latest work  updates
	go func() {
		// watch
		workCh := ma.etcdCli.Watch(context.Background(), wemixWorkKey)
		lockCh := ma.etcdCli.Watch(context.Background(), wemixTokenKey)
		for {
			if !ma.etcdIsRunning() {
				break
			}
			select {
			case <-etcd.Server.LeaderChangedNotify():
				latestEtcdLeader.Store(ma.etcd.Server.Leader())
			case watchResp := <-workCh:
				latestUpdateTime.Store(time.Now())
				for _, event := range watchResp.Events {
					switch event.Type {
					case mvccpb.PUT:
						latestWemixWork.Store(event.Kv.Value)
					case mvccpb.DELETE:
						var nilBytes []byte
						latestWemixWork.Store(nilBytes)
					}
				}
			case watchResp := <-lockCh:
				for _, event := range watchResp.Events {
					switch event.Type {
					case mvccpb.PUT:
						latestMiningToken.Store(event.Kv.Value)
					case mvccpb.DELETE:
						var nilBytes []byte
						latestMiningToken.Store(nilBytes)
					}
				}
			}
		}
	}()
	return nil
}

func (ma *wemixAdmin) etcdJoin_old(cluster string) error {
	if ma.etcdIsRunning() {
		return ErrAlreadyRunning
	}

	cfg := ma.etcdNewConfig(false)
	cfg.InitialCluster = cluster
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Error("failed to join etcd", "error", err)
		return err
	} else {
		log.Info("started etcd server")
	}

	ma.etcd = etcd
	ma.etcdCli = v3client.New(etcd.Server)
	return nil
}

func (ma *wemixAdmin) etcdJoin(name string) error {
	var node *wemixNode

	ma.lock.Lock()
	for _, i := range ma.nodes {
		if i.Name == name || i.Enode == name || i.Id == name || i.Ip == name {
			node = i
			break
		}
	}
	ma.lock.Unlock()

	if node == nil {
		return ethereum.NotFound
	}

	ch := make(chan string, 16)
	sub := wemixapi.SubscribeToEtcdCluster(ch)
	defer func() {
		sub.Unsubscribe()
		close(ch)
	}()

	timer := time.NewTimer(30 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	err := admin.rpcCli.CallContext(ctx, nil, "admin_requestEtcdAddMember", &node.Id)
	cancel()
	if err != nil {
		log.Error("admin_requestEtcdAddMember failed", "id", node.Id, "error", err)
		return err
	}

	for {
		select {
		case cluster := <-ch:
			cluster, _ = ma.etcdFixCluster(cluster)

			cfg := ma.etcdNewConfig(false)
			cfg.InitialCluster = cluster
			etcd, err := embed.StartEtcd(cfg)
			if err != nil {
				log.Error("failed to join etcd", "error", err)
				return err
			} else {
				log.Info("started etcd server")
			}

			ma.etcd = etcd
			ma.etcdCli = v3client.New(etcd.Server)
			return nil

		case <-timer.C:
			return fmt.Errorf("Timed Out")
		}
	}
}

func (ma *wemixAdmin) etcdStop() error {
	if !ma.etcdIsRunning() {
		return ErrNotRunning
	}
	if ma.etcdCli != nil {
		ma.etcdCli.Close()
	}
	if ma.etcd != nil {
		ma.etcd.Server.HardStop()
	}
	ma.etcd = nil
	ma.etcdCli = nil
	return nil
}

func (ma *wemixAdmin) etcdIsLeader() bool {
	if !ma.etcdIsRunning() {
		return false
	} else {
		return ma.etcd.Server.ID() == ma.etcd.Server.Leader()
	}
}

// returns leader id and node
func (ma *wemixAdmin) etcdLeader(locked bool) (uint64, *wemixNode) {
	if !ma.etcdIsRunning() {
		return 0, nil
	}

	lid := uint64(ma.etcd.Server.Leader())
	for _, i := range ma.etcd.Server.Cluster().Members() {
		if uint64(i.ID) == lid {
			var node *wemixNode
			if !locked {
				ma.lock.Lock()
			}
			for _, j := range ma.nodes {
				if i.Attributes.Name == j.Name {
					node = j
					break
				}
			}
			if !locked {
				ma.lock.Unlock()
			}
			return lid, node
		}
	}

	return 0, nil
}

func (ma *wemixAdmin) etcdPut(key, value string) (int64, error) {
	if !ma.etcdIsRunning() {
		return 0, ErrNotRunning
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		ma.etcd.Server.Cfg.ReqTimeout())
	defer cancel()
	resp, err := ma.etcdCli.Put(ctx, key, value)
	if err == nil {
		return resp.Header.Revision, err
	} else {
		return 0, err
	}
}

func (ma *wemixAdmin) etcdGet(key string) (string, error) {
	if !ma.etcdIsRunning() {
		return "", ErrNotRunning
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(1)*time.Second)
	defer cancel()
	rsp, err := ma.etcdCli.Get(ctx, key)
	if err != nil {
		return "", err
	} else if rsp.Count == 0 {
		return "", ErrNotFound
	} else {
		var v string
		for _, kv := range rsp.Kvs {
			v = string(kv.Value)
		}
		return v, nil
	}
}

// compare & swap, do put only if previous value matches
func (ma *wemixAdmin) etcdPut2(key, value, prev string) error {
	if !ma.etcdIsRunning() {
		return ErrNotRunning
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		ma.etcd.Server.Cfg.ReqTimeout())
	defer cancel()

	tx := ma.etcdCli.Txn(ctx)
	txresp, err := tx.If(
		clientv3.Compare(clientv3.Value(key), "=", prev),
	).Then(
		clientv3.OpPut(key, value),
		clientv3.OpGet(key),
	).Else(
		clientv3.OpGet(key),
	).Commit()
	if err == nil && !txresp.Succeeded {
		err = ErrExists
	}
	return err
}

func (ma *wemixAdmin) etcdDelete(key string) error {
	if !ma.etcdIsRunning() {
		return ErrNotRunning
	}
	ctx, cancel := context.WithTimeout(context.Background(),
		ma.etcd.Server.Cfg.ReqTimeout())
	defer cancel()
	_, err := ma.etcdCli.Delete(ctx, key)
	return err
}

// handles removed nodes
// caller should take care of etcd & governance lock
func etcdSyncMembership() error {
	if admin == nil || !admin.amPartner() || admin.self == nil || !admin.etcdIsRunning() {
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	header, err := admin.cli.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Error("etcd sync: failed to get the latest block", "error", err)
		return err
	}
	nodes, err := getNodesAt(header.Number)
	if err != nil {
		return err
	}
	m := map[string]*wemixNode{}
	for _, n := range nodes {
		m[n.Name] = n
	}
	members := admin.etcd.Server.Cluster().Members()
	if len(nodes) > len(members) {
		// some nodes are not added to etcd yet
		return nil
	}

	var removed []*membership.Member
	for _, cn := range members {
		if _, ok := m[cn.Attributes.Name]; !ok {
			removed = append(removed, cn)
		}
	}
	if len(removed) == 0 {
		return nil
	}

	// remove the first one
	_, err = admin.etcdRemoveMember(removed[0].Attributes.Name)
	log.Error("etcd node sync", "removed-node", removed[0].Attributes.Name, "error", err)
	return err
}

// leases

type WemixToken struct {
	admin  *wemixAdmin
	Miner  string `json:"miner"`
	ID     uint64 `json:"id"`
	Height int64  `json:"height"`
	Since  int64  `json:"since"`
	Till   int64  `json:"till"`
	Key    string `json:"key"`
}

func (ma *wemixAdmin) acquireToken(ctx context.Context, height *big.Int, ttl int) (*WemixToken, error) {
again:
	now := time.Now().Unix()
	till := now + int64(ttl)
	lock := &WemixToken{
		admin:  ma,
		Miner:  ma.self.Name,
		ID:     uint64(ma.etcd.Server.ID()),
		Height: height.Int64(),
		Since:  now,
		Till:   till,
		Key:    wemixTokenKey,
	}
	value, err := json.Marshal(lock)
	if err != nil {
		return nil, err
	}

	tx := ma.etcdCli.Txn(ctx)
	txresp, err := tx.If(
		clientv3.Compare(clientv3.CreateRevision(wemixTokenKey), "=", 0),
	).Then(
		clientv3.OpPut(wemixTokenKey, string(value)),
	).Else(
		clientv3.OpGet(wemixTokenKey),
	).Commit()

	if err == nil && !txresp.Succeeded {
		err = ErrExists
		var (
			tokenFound bool = false
			foundToken []byte
		)
		for _, r := range txresp.Responses {
			for _, kv := range r.GetResponseRange().Kvs {
				switch string(kv.Key) {
				case wemixTokenKey:
					tokenFound = true
					foundToken = kv.Value
				}
			}
		}

		if tokenFound {
			if len(foundToken) > 0 {
				var otherToken = &WemixToken{}
				if err = json.Unmarshal(foundToken, otherToken); err != nil {
					return nil, err
				}
				otherToken.admin = ma
				if otherToken.Till >= time.Now().Unix() {
					// valid lock
					return otherToken, ErrExists
				}
			}

			// expired or empty lock, delete it & try again
			tx = ma.etcdCli.Txn(ctx)
			_, err := tx.If(
				clientv3.Compare(clientv3.Value(wemixTokenKey), "=", string(foundToken)),
			).Then(
				clientv3.OpDelete(wemixTokenKey),
			).Commit()
			if err != nil {
				return nil, err
			}
			goto again
		}
	}
	return lock, err
}

func (lck *WemixToken) ttl() int64 {
	ttl := lck.Till - time.Now().Unix()
	if ttl < 0 {
		ttl = -1
	}
	return ttl
}

func (lck *WemixToken) renew(ctx context.Context, ttl int) error {
	prev, err := json.Marshal(lck)
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	prevTill := lck.Till
	lck.Till = now + int64(ttl)
	value, err := json.Marshal(lck)
	if err != nil {
		return err
	}

	tx := lck.admin.etcdCli.Txn(ctx)
	txresp, err := tx.If(
		clientv3.Compare(clientv3.Value(lck.Key), "=", prev),
	).Then(
		clientv3.OpPut(lck.Key, string(value)),
		clientv3.OpGet(lck.Key),
	).Else(
		clientv3.OpGet(lck.Key),
	).Commit()
	if err == nil && !txresp.Succeeded {
		err = ErrExists
	}
	if err != nil {
		// restore the till value upon failure
		lck.Till = prevTill
	}
	return err
}

func (lck *WemixToken) release(ctx context.Context) error {
	value, err := json.Marshal(lck)
	if err != nil {
		return err
	}

	tx := lck.admin.etcdCli.Txn(ctx)
	txresp, err := tx.If(
		clientv3.Compare(clientv3.Value(lck.Key), "=", string(value)),
	).Then(
		clientv3.OpDelete(lck.Key),
	).Commit()

	if err == nil && !txresp.Succeeded {
		return ErrExists
	}
	return err
}

func (lck *WemixToken) lockedPut(ctx context.Context, key, value, prev string) error {
	exists := true
	lockValue, err := json.Marshal(lck)
	if err != nil {
		return err
	}

again:
	tx := lck.admin.etcdCli.Txn(ctx)
	var txIf clientv3.Txn
	if exists {
		txIf = tx.If(
			clientv3.Compare(clientv3.Value(lck.Key), "=", lockValue),
			clientv3.Compare(clientv3.Value(key), "=", prev),
		)
	} else {
		txIf = tx.If(
			clientv3.Compare(clientv3.Value(lck.Key), "=", lockValue),
			clientv3.Compare(clientv3.Version(key), "=", 0),
		)
	}
	txresp, err := txIf.Then(
		clientv3.OpPut(key, string(value)),
		clientv3.OpGet(key),
	).Else(
		clientv3.OpGet(key),
	).Commit()

	if err == nil && !txresp.Succeeded {
		err = ErrExists
		if len(txresp.Responses) > 0 {
			if rr := txresp.Responses[0].GetResponseRange(); rr.Count == 0 {
				// if not exists, put empty string and try again
				tx = lck.admin.etcdCli.Txn(ctx)
				_, err := tx.If(
					clientv3.Compare(clientv3.Version(key), "=", 0),
				).Then(
					clientv3.OpPut(key, ""),
				).Commit()
				if err == nil {
					exists = false
					goto again
				}
			}
		}

	}
	return err
}

func (ma *wemixAdmin) ttl2(ctx context.Context, key string) (int64, error) {
	rsp, err := ma.etcdCli.Get(ctx, key)
	if err != nil {
		return -1, err
	} else if rsp.Count == 0 {
		return -1, nil
	}
	lock := &WemixToken{}
	var value []byte
	for _, kv := range rsp.Kvs {
		value = kv.Value
		break
	}
	if e2 := json.Unmarshal(value, lock); e2 != nil {
		err = e2
		return -1, err
	}
	lock.admin = ma

	ttl := lock.Till - time.Now().Unix()
	if ttl < 0 {
		ttl = -1
	}
	return ttl, nil
}

// acquire token iff we're in sync with the latest block
// lock is expected not to be present
//   if expired, delete & try again
// work is expected to be there
//   if not present, put an empty string & try again
func (ma *wemixAdmin) acquireTokenSync(ctx context.Context, height *big.Int, parentHash common.Hash, ttl int64) (*WemixToken, error) {
	if ok, err := ma.isEligibleMiner(height); err != nil {
		return nil, err
	} else if !ok {
		return nil, ErrIneligible
	}

	workExists := true
	prevWork := ""
	if prevWorkBytes, err := json.Marshal(&wemixWork{Height: height.Int64() - 1, Hash: parentHash}); err != nil {
		return nil, err
	} else {
		prevWork = string(prevWorkBytes)
	}

again:
	now := time.Now().Unix()
	till := now + ttl
	lock := &WemixToken{
		admin:  ma,
		Miner:  ma.self.Name,
		ID:     uint64(ma.etcd.Server.ID()),
		Height: height.Int64(),
		Since:  now,
		Till:   till,
		Key:    wemixTokenKey,
	}
	lockValue, err := json.Marshal(lock)
	if err != nil {
		return nil, err
	}

	tx := ma.etcdCli.Txn(ctx)
	var txIf clientv3.Txn
	if workExists {
		txIf = tx.If(
			clientv3.Compare(clientv3.CreateRevision(wemixTokenKey), "=", 0),
			clientv3.Compare(clientv3.Value(wemixWorkKey), "=", prevWork),
		)
	} else {
		txIf = tx.If(
			clientv3.Compare(clientv3.CreateRevision(wemixTokenKey), "=", 0),
			clientv3.Compare(clientv3.CreateRevision(wemixWorkKey), "=", 0),
		)
	}
	txresp, err := txIf.Then(
		clientv3.OpPut(wemixTokenKey, string(lockValue)),
	).Else(
		clientv3.OpGet(wemixTokenKey),
		clientv3.OpGet(wemixWorkKey),
	).Commit()

	if err == nil && !txresp.Succeeded {
		err = ErrExists
		var (
			tokenFound, workFound bool = false, false
			foundToken            []byte
			foundWork             []byte
		)
		for _, r := range txresp.Responses {
			for _, kv := range r.GetResponseRange().Kvs {
				switch string(kv.Key) {
				case wemixTokenKey:
					tokenFound = true
					foundToken = kv.Value
				case wemixWorkKey:
					workFound = true
					foundWork = kv.Value
				}
			}
		}

		if tokenFound {
			if len(foundToken) > 0 {
				var otherToken = &WemixToken{}
				if err = json.Unmarshal(foundToken, otherToken); err != nil {
					return nil, err
				}
				otherToken.admin = ma
				if otherToken.Till >= time.Now().Unix() {
					// valid lock
					return otherToken, ErrExists
				}
			}

			// expired or empty lock, delete it & try again
			tx = ma.etcdCli.Txn(ctx)
			_, err := tx.If(
				clientv3.Compare(clientv3.Value(wemixTokenKey), "=", string(foundToken)),
			).Then(
				clientv3.OpDelete(wemixTokenKey),
			).Commit()
			if err != nil {
				return nil, err
			}
			goto again
		}

		// lock not found, but work doesn't match
		if workFound {
			_ = foundWork
			return nil, ErrInvalidWork
		} else if workExists {
			workExists = false
			goto again
		}
		return nil, ErrInvalidWork
	}
	if lock != nil && err == nil {
		etcdSyncMembership()
	}
	return lock, err
}

func (lck *WemixToken) releaseTokenSync(ctx context.Context, height *big.Int, hash, parentHash common.Hash) error {
	exists := true
	prevWork, work, lockValue := "", "", ""

	if data, err := json.Marshal(lck); err != nil {
		return err
	} else {
		lockValue = string(data)
	}
	if data, err := json.Marshal(&wemixWork{Height: height.Int64() - 1, Hash: parentHash}); err != nil {
		return err
	} else {
		prevWork = string(data)
	}
	if data, err := json.Marshal(&wemixWork{Height: height.Int64(), Hash: hash}); err != nil {
		return err
	} else {
		work = string(data)
	}

again:
	tx := lck.admin.etcdCli.Txn(ctx)
	var txIf clientv3.Txn
	if exists {
		txIf = tx.If(
			clientv3.Compare(clientv3.Value(lck.Key), "=", lockValue),
			clientv3.Compare(clientv3.Value(wemixWorkKey), "=", prevWork),
		)
	} else {
		txIf = tx.If(
			clientv3.Compare(clientv3.Value(lck.Key), "=", lockValue),
			clientv3.Compare(clientv3.Version(wemixWorkKey), "=", 0),
		)
	}
	txresp, err := txIf.Then(
		clientv3.OpDelete(lck.Key),
		clientv3.OpPut(wemixWorkKey, work),
	).Else(
		clientv3.OpGet(lck.Key),
		clientv3.OpGet(wemixWorkKey),
	).Commit()

	if err == nil && !txresp.Succeeded {
		var (
			tokenFound, workFound bool = false, false
			foundToken            []byte
		)
		for _, r := range txresp.Responses {
			for _, kv := range r.GetResponseRange().Kvs {
				switch string(kv.Key) {
				case wemixTokenKey:
					tokenFound = true
					foundToken = kv.Value
				case wemixWorkKey:
					workFound = true
				}
			}
		}

		if !tokenFound || lockValue != string(foundToken) {
			// we don't have the lock
			return ErrInvalidToken
		}
		if !workFound {
			exists = false
			goto again
		}
		// work mismatch
		err = ErrInvalidWork
	}
	return err
}

func (ma *wemixAdmin) etcdCompact(rev int64) error {
	if !ma.etcdIsRunning() {
		return ErrNotRunning
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		ma.etcd.Server.Cfg.ReqTimeout())
	defer cancel()
	_, err := ma.etcdCli.Compact(ctx, rev, clientv3.WithCompactPhysical())
	// WithCompactPhysical makes Compact wait until all compacted entries are
	// removed from the etcd server's storage.
	return err
}

func (ma *wemixAdmin) etcdInfo() interface{} {
	if ma.etcd == nil {
		return ErrNotRunning
	}

	getMemberInfo := func(member *etcdserverpb.Member) *map[string]interface{} {
		return &map[string]interface{}{
			"name":       member.Name,
			"id":         fmt.Sprintf("%x", member.ID),
			"clientUrls": strings.Join(member.ClientURLs, ","),
			"peerUrls":   strings.Join(member.PeerURLs, ","),
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		ma.etcd.Server.Cfg.ReqTimeout())
	rsp, err := ma.etcdCli.MemberList(ctx)
	cancel()

	var ms []*etcdserverpb.Member
	if err == nil {
		ms = append(ms, rsp.Members...)
		sort.Slice(ms, func(i, j int) bool {
			return ms[i].Name < ms[j].Name
		})
	}

	var bb bytes.Buffer
	var self, leader *etcdserverpb.Member
	var members []interface{}
	for _, i := range ms {
		if i.ID == uint64(ma.etcd.Server.ID()) {
			self = i
		}
		if i.ID == uint64(ma.etcd.Server.Leader()) {
			leader = i
		}
		members = append(members, getMemberInfo(i))
		if bb.Len() > 0 {
			bb.WriteString(",")
		}
		bb.WriteString(fmt.Sprintf("%s=%s", i.Name,
			strings.Join(i.PeerURLs, ",")))
	}

	info := map[string]interface{}{
		"cluster": bb.String(),
		"members": members,
	}
	if self != nil {
		info["self"] = &map[string]interface{}{
			"name": self.Name,
			"id":   fmt.Sprintf("%x", self.ID),
		}
	}
	if leader != nil {
		info["leader"] = &map[string]interface{}{
			"name": leader.Name,
			"id":   fmt.Sprintf("%x", leader.ID),
		}
	}

	return info
}

func EtcdInit() error {
	etcdLock.Lock()
	defer etcdLock.Unlock()

	if admin == nil {
		return ErrNotRunning
	}
	return admin.etcdInit()
}

func EtcdStart() {
	if !etcdLock.TryLock() {
		return
	}
	defer etcdLock.Unlock()
	if admin == nil {
		return
	}

	admin.etcdStart()
	if !admin.etcdIsRunning() {
		// try to join a random peer
		var node *wemixNode
		admin.lock.Lock()
		if len(admin.nodes) > 0 {
			ix := rand.Int() % len(admin.nodes)
			for _, i := range admin.nodes {
				if ix <= 0 {
					node = i
					break
				}
				ix--
			}
		}
		admin.lock.Unlock()

		if node != nil && admin.isPeerUp(node.Id) {
			log.Info("Wemix", "Trying to join", node.Name)
			admin.etcdJoin(node.Name)
		}
	}
}

func EtcdAddMember(name string) (string, error) {
	etcdLock.Lock()
	defer etcdLock.Unlock()

	if admin == nil {
		return "", ErrNotRunning
	}
	return admin.etcdAddMember(name)
}

func EtcdRemoveMember(name string) (string, error) {
	etcdLock.Lock()
	defer etcdLock.Unlock()

	if admin == nil {
		return "", ErrNotRunning
	}
	return admin.etcdRemoveMember(name)
}

func EtcdMoveLeader(name string) error {
	etcdLock.Lock()
	defer etcdLock.Unlock()

	if admin == nil {
		return ErrNotRunning
	}
	return admin.etcdMoveLeader(name)
}

func EtcdJoin(name string) error {
	etcdLock.Lock()
	defer etcdLock.Unlock()

	if admin == nil {
		return ErrNotRunning
	}
	return admin.etcdJoin(name)
}

func EtcdGetWork() (string, error) {
	if admin == nil {
		return "", ErrNotRunning
	}
	return admin.etcdGet(wemixWorkKey)
}

func EtcdDeleteWork() error {
	if admin == nil {
		return ErrNotRunning
	}
	return admin.etcdDelete(wemixWorkKey)
}

func EtcdPut(key, value string) error {
	if admin == nil {
		return ErrNotRunning
	}
	_, err := admin.etcdPut(key, value)
	return err
}

func EtcdGet(key string) (string, error) {
	if admin == nil {
		return "", ErrNotRunning
	}
	return admin.etcdGet(key)
}

func EtcdDelete(key string) error {
	if admin == nil {
		return ErrNotRunning
	}
	return admin.etcdDelete(key)
}

/* EOF */
