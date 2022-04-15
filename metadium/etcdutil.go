/* etcdutil.go */

package metadium

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"go.etcd.io/etcd/server/v3/embed"
	"go.etcd.io/etcd/server/v3/etcdserver/api/membership"
	"go.etcd.io/etcd/server/v3/etcdserver/api/v3client"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/log"
	metaapi "github.com/ethereum/go-ethereum/metadium/api"
)

var (
	etcdLock = &SpinLock{0}
)

func (ma *metaAdmin) etcdMemberExists(name, cluster string) (bool, error) {
	var node *metaNode
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
// "=http://1.1.1.1:8590,meta2=http:/1.1.1.2:8590"
func (ma *metaAdmin) etcdFixCluster(cluster string) (string, error) {
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

func (ma *metaAdmin) etcdNewConfig(newCluster bool) *embed.Config {
	// LPUrls: listening peer urls
	// APUrls: advertised peer urls
	// LCUrls: listening client urls
	// LPUrls: advertised client urls
	cfg := embed.NewConfig()
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

func (ma *metaAdmin) etcdIsRunning() bool {
	return ma.etcd != nil && ma.etcdCli != nil
}

func (ma *metaAdmin) etcdGetCluster() string {
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
func (ma *metaAdmin) etcdAddMember(name string) (string, error) {
	if !ma.etcdIsRunning() {
		return "", ErrNotRunning
	}

	if ok, _ := ma.etcdMemberExists(name, ma.etcdGetCluster()); ok {
		return ma.etcdGetCluster(), nil
	}

	var node *metaNode
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
		log.Error("Metadium: failed to add a new member",
			"name", name, "ip", node.Ip, "port", node.Port+1, "error", err)
		return "", err
	} else {
		log.Info("Metadium: a new member added",
			"name", name, "ip", node.Ip, "port", node.Port+1, "error", err)
		return ma.etcdGetCluster(), nil
	}
}

// returns new cluster string if removing the member is successful
func (ma *metaAdmin) etcdRemoveMember(name string) (string, error) {
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

func (ma *metaAdmin) etcdMoveLeader(name string) error {
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

	ctx, cancel := context.WithTimeout(context.Background(), ma.etcd.Server.Cfg.ReqTimeout())
	err := ma.etcd.Server.MoveLeader(ctx, ma.etcd.Server.Lead(), id)
	defer cancel()
	return err
}

func (ma *metaAdmin) etcdWipe() error {
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

func (ma *metaAdmin) etcdInit() error {
	if ma.etcdIsRunning() {
		return ErrAlreadyRunning
	} else if ma.self == nil {
		return ErrNotRunning
	}

	cfg := ma.etcdNewConfig(true)
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Error("Metadium: failed to initialize etcd", "error", err)
		return err
	} else {
		log.Info("Metadium: initialized etcd server")
	}

	ma.etcd = etcd
	ma.etcdCli = v3client.New(etcd.Server)
	return nil
}

func (ma *metaAdmin) etcdStart() error {
	if ma.etcdIsRunning() {
		return ErrAlreadyRunning
	}

	cfg := ma.etcdNewConfig(false)
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Error("Metadium: failed to start etcd", "error", err)
		return err
	} else {
		log.Info("Metadium: started etcd server")
	}

	ma.etcd = etcd
	ma.etcdCli = v3client.New(etcd.Server)
	return nil
}

func (ma *metaAdmin) etcdJoin_old(cluster string) error {
	if ma.etcdIsRunning() {
		return ErrAlreadyRunning
	}

	cfg := ma.etcdNewConfig(false)
	cfg.InitialCluster = cluster
	etcd, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Error("Metadium: failed to join etcd", "error", err)
		return err
	} else {
		log.Info("Metadium: started etcd server")
	}

	ma.etcd = etcd
	ma.etcdCli = v3client.New(etcd.Server)
	return nil
}

func (ma *metaAdmin) etcdJoin(name string) error {
	var node *metaNode

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

	msgch := make(chan interface{}, 32)
	metaapi.SetMsgChannel(msgch)
	defer func() {
		metaapi.SetMsgChannel(nil)
		close(msgch)
	}()

	timer := time.NewTimer(30 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	err := admin.rpcCli.CallContext(ctx, nil, "admin_requestEtcdAddMember", &node.Id)
	cancel()
	if err != nil {
		log.Error("Metadium admin_requestEtcdAddMember failed", "id", node.Id, "error", err)
		return err
	}

	for {
		select {
		case msg := <-msgch:
			cluster, ok := msg.(string)
			if !ok {
				continue
			}

			cluster, _ = ma.etcdFixCluster(cluster)

			cfg := ma.etcdNewConfig(false)
			cfg.InitialCluster = cluster
			etcd, err := embed.StartEtcd(cfg)
			if err != nil {
				log.Error("Metadium: failed to join etcd", "error", err)
				return err
			} else {
				log.Info("Metadium: started etcd server")
			}

			ma.etcd = etcd
			ma.etcdCli = v3client.New(etcd.Server)
			return nil

		case <-timer.C:
			return fmt.Errorf("Timed Out")
		}
	}
}

func (ma *metaAdmin) etcdStop() error {
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

func (ma *metaAdmin) etcdIsLeader() bool {
	if !ma.etcdIsRunning() {
		return false
	} else {
		return ma.etcd.Server.ID() == ma.etcd.Server.Leader()
	}
}

// returns leader id and node
func (ma *metaAdmin) etcdLeader(locked bool) (uint64, *metaNode) {
	if !ma.etcdIsRunning() {
		return 0, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		ma.etcd.Server.Cfg.ReqTimeout())
	rsp, err := ma.etcdCli.MemberList(ctx)
	cancel()

	if err != nil {
		return 0, nil
	}

	lid := uint64(ma.etcd.Server.Leader())
	for _, i := range rsp.Members {
		if i.ID == lid {
			var node *metaNode
			if !locked {
				ma.lock.Lock()
			}
			for _, j := range ma.nodes {
				if i.Name == j.Name {
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

func (ma *metaAdmin) etcdPut(key, value string) error {
	if !ma.etcdIsRunning() {
		return ErrNotRunning
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		ma.etcd.Server.Cfg.ReqTimeout())
	defer cancel()
	_, err := ma.etcdCli.Put(ctx, key, value)
	return err
}

func (ma *metaAdmin) etcdGet(key string) (string, error) {
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
		return "", nil
	} else {
		var v string
		for _, kv := range rsp.Kvs {
			v = string(kv.Value)
		}
		return v, nil
	}
}

func (ma *metaAdmin) etcdDelete(key string) error {
	if !ma.etcdIsRunning() {
		return ErrNotRunning
	}
	ctx, cancel := context.WithTimeout(context.Background(),
		ma.etcd.Server.Cfg.ReqTimeout())
	defer cancel()
	_, err := ma.etcdCli.Delete(ctx, key)
	return err
}

func (ma *metaAdmin) etcdInfo() interface{} {
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
		var node *metaNode
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
			log.Info("Metadium", "Trying to join", node.Name)
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
	return admin.etcdGet("metadium-work")
}

func EtcdDeleteWork() error {
	if admin == nil {
		return ErrNotRunning
	}
	return admin.etcdDelete("metadium-work")
}

/* EOF */
