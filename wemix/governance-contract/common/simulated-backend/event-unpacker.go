package sim

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/wemix/governance-contract/common/bn"
)

/*
- 전체로그 검색
sim.NewEventLogsUnpacker(receipt).FindMany(nil, nil, nil)
- 특정 컨트랙트 출력
sim.NewEventLogsUnpacker(receipt).FindMany(contract, nil, nil)
- 특정 컨트랙트의 이벤트 검색
sim.NewEventLogsUnpacker(receipt).FindMany(contract, eventID, nil)
- 특정 컨트랙트의 이벤트 검색
sim.NewEventLogsUnpacker(receipt).FindMany(contract, eventID, nil)
- 특정 이벤트 모두 검색
sim.NewEventLogsUnpacker(receipt).FindMany(nil, eventID, nil)
- filter 검색
sim.NewEventLogsUnpacker(receipt).FindOne(nil, eventID, sim.Arguments{
			"from":  common.Address{},
			"to":    address,
			"value": amount,
		}))
- 출력
t.Log(sim.NewEventLogsUnpacker(receipt).FindMany(nil, nil, nil))
*/

type allContractsType map[common.Address]string
type allEventsType map[common.Hash][]abi.Event

func (a allEventsType) String() string {
	result := fmt.Sprintf("all events type: %v", len(a))
	total := 0

	for hash, es := range a {
		for _, e := range es {
			result = fmt.Sprintf("%s\n%v: %v", result, hash, e.String())
			total++
		}
	}
	result = fmt.Sprintf("%s\ntotal all events loaded %v", result, total)
	return result
}

var (
	allEventsLock = sync.RWMutex{}
	allEvents     = allEventsType{}    // eventID => abi.Event
	allContracts  = allContractsType{} // CA => contract.Alias

	// contract생성시, EOA 생성시 address를 alias와 매칭시켜, string으로 변환시 해당 address를 alias로 대체한다.
	// string으로 변환시 zero address는 0x0변경은 디폴트로 설정
	replacer = []string{
		common.Address{}.Hex(), "0x00",
		bn.MaxUint256.String(), "uint256_max",
		bn.Sub(bn.MaxUint256, 1).String(), "uint256_max-1",
	}
)

// event를 모음.
func collectEvents(events map[string]abi.Event) error {
	for _, e := range events {
		func() {
			allEventsLock.Lock()
			defer allEventsLock.Unlock()
			for _, exist := range allEvents[e.ID] {
				// 이벤트 full name으로 비교하여 이미 있는 이름은 넘어감.
				if exist.String() == e.String() {
					goto skip
				}
			}
			allEvents[e.ID] = append(allEvents[e.ID], e)
		skip:
		}()
	}
	return nil
}

// contract 모음.
func collectContract(ca common.Address, alias string) {
	allContracts[ca] = alias
	replacer = append(replacer, hexutil.Encode(ca[:]), alias)
}

type EOA struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

var EOAKey = map[common.Address]*ecdsa.PrivateKey{}
var EOAAlias = map[common.Address]string{}
var AliasToAddress = map[string]common.Address{}

// EOA 생성 및 alias 등록
// event print시 해당 address는 alias로 표시
func NewEOA(alias string) (eoa *EOA) {
	pk, _ := crypto.GenerateKey()
	return NewPrivKey(alias, pk)
}

func NewPrivKey(alias string, privKey *ecdsa.PrivateKey) (eoa *EOA) {
	eoa = &EOA{
		PrivateKey: privKey,
		Address:    crypto.PubkeyToAddress(privKey.PublicKey),
	}
	EOAKey[eoa.Address] = eoa.PrivateKey
	EOAAlias[eoa.Address] = alias
	if before, ok := AliasToAddress[alias]; ok {
		fmt.Printf("%s is over writted (%v -> %v)", alias, before, eoa.Address)
	}
	AliasToAddress[alias] = eoa.Address

	replacer = append(replacer, hexutil.Encode(eoa.Address[:]), alias)
	return
}

func GetOrNewEOA(alias string) common.Address {
	addr, ok := AliasToAddress[alias]
	if !ok {
		addr = NewEOA(alias).Address
	}
	return addr
}

// event arguments대한 변수명 : 값 타입
type Arguments map[string]interface{}

type UnpackedOutput struct {
	ContractAlias string
	Contract      common.Address
	EventName     string
	Event         common.Hash
	EventIndex    uint      // tx상의 event index
	Arguments     Arguments // argument
}

// address중 alias가 있다면 별칭으로 바꾼다.
func (u *UnpackedOutput) String() string {
	arguments, _ := json.MarshalIndent(u.Arguments, "", "  ")
	// replacer에 등록된 address를 alias로 변환
	argumentsReplaced := strings.NewReplacer(replacer...).Replace(string(arguments))

	return fmt.Sprintf("%d [%s] %s%v", u.EventIndex, u.ContractAlias, u.EventName, argumentsReplaced)
}

// event 하나에대한 타입
type Unpacked struct {
	ContractAlias  string
	Contract       common.Address
	EventName      string
	Event          common.Hash
	EventIndex     uint        // tx상의 event index
	ArgumentsSlice []Arguments // argument
}

type UnpackedArgumentsOutput []*UnpackedOutput

func (u UnpackedArgumentsOutput) String() string {
	result := fmt.Sprintf("filtered logs, total logs : %v", len(u))
	for _, unpacked := range u {
		result = fmt.Sprintf("%s\n%s", result, unpacked)
	}
	return result
}

type UnpackedArguments []*Unpacked

// EventLogsUnpacker : tx receipt의 모든 log를 unpack 및 search
type EventLogsUnpacker struct {
	Receipt *types.Receipt
	All     UnpackedArguments
}

func NewEventLogsUnpacker(receipt *types.Receipt) *EventLogsUnpacker {
	u := &EventLogsUnpacker{
		Receipt: receipt,
		All:     make(UnpackedArguments, 0),
	}
	u.unpack()
	return u
}

func (u *EventLogsUnpacker) unpackLog(e *types.Log) (argumentsSlice []Arguments, event *abi.Event) {
	argumentsSlice = []Arguments{}

	allEventsLock.RLock()
	defer allEventsLock.RUnlock()

	for _, fn := range allEvents[e.Topics[0]] {
		arguments := Arguments{}

		// topics 또는 data로 부터 unpack할 것을 구분
		var indexed, nonIndexed abi.Arguments
		for _, arg := range fn.Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			} else {
				nonIndexed = append(nonIndexed, arg)
			}
		}

		// 만약 data size가 0이면 이벤트타입의 input도 0이여야한다.
		if len(e.Data) == 0 {
			if len(nonIndexed) > 0 {
				continue
			}
		} else {
			// data를 unpack
			if marshalledValues, err := fn.Inputs.UnpackValues(e.Data); err != nil {
				continue
			} else {
				for i, arg := range nonIndexed {
					arguments[arg.Name] = marshalledValues[i]
				}
			}
		}

		// topic을 unpack
		if len(e.Topics[1:]) != len(indexed) {
			continue
		} else if err := abi.ParseTopicsIntoMap(arguments, indexed, e.Topics[1:]); err != nil {
			continue
		}
		event = &fn
		// hash가 byte array로 변환이 되니 변환한다.
		for k, arg := range arguments {
			switch v := arg.(type) {
			case [32]uint8:
				arguments[k] = common.BytesToHash(v[:])
			}
		}
		argumentsSlice = append(argumentsSlice, arguments)

	}
	return
}

// receipt의 모든 log를 unpack
func (u *EventLogsUnpacker) unpack() {
	for _, e := range u.Receipt.Logs {
		argumentsSlice, event := u.unpackLog(e)
		if event == nil {
			continue
		}
		// 배포한 모든 컨트랙트에서 어드레스로 별칭을 가져온다.
		alias, ok := allContracts[e.Address]
		if !ok {
			// 만약 없다면 address를 별칭으로 넣는다.
			alias = common.Bytes2Hex(e.Address[:])
		}

		// 완료
		u.All = append(u.All, &Unpacked{
			ContractAlias:  alias,
			Contract:       e.Address,
			EventName:      event.Name,
			Event:          e.Topics[0],
			EventIndex:     e.Index,
			ArgumentsSlice: argumentsSlice,
		})
	}
}

// 조건에 맞는 모든 log를 검색하여 반환
// ca == nil 이면 모든 컨트랙트, 기본타입은 common.Address
// eventID == nil 이면 모든 event id, 기본타입은 common.Hash
// filter는 log arguments에대한 조건
func (u *EventLogsUnpacker) FindMany(ca interface{}, eventID interface{}, filter Arguments) (events UnpackedArgumentsOutput) {
	// ca와 eventID를 common.Address, common.Hash타입으로 변환을하고,
	// 0x00이 아니면 먼저 bloom필터로 먼저 검사후 진행한다.
	if ca = toAddress(ca); ca != (common.Address{}) {
		if !u.Receipt.Bloom.Test(ca.(common.Address).Bytes()) {
			return
		}
	}
	if eventID = toHash(eventID); eventID != (common.Hash{}) {
		if !u.Receipt.Bloom.Test(eventID.(common.Hash).Bytes()) {
			return
		}
	}

	events = make(UnpackedArgumentsOutput, 0)
	for _, e := range u.All {
		// 컨트랙트 조건 필터링
		if ca != (common.Address{}) && ca != e.Contract {
			continue
		}
		// 이벤트 조건 필터링
		if eventID != (common.Hash{}) && eventID != e.Event {
			continue
		}
	loop:
		for _, arguments := range e.ArgumentsSlice {
			// 입력한 필터의 값을 필터링
			for k, v := range filter {
				if a := arguments[k]; a == nil {
					goto skip
				} else if !reflect.DeepEqual(a, v) {
					goto skip
				}
			}
			// 조건통과
			events = append(events, &UnpackedOutput{
				ContractAlias: e.ContractAlias,
				Contract:      e.Contract,
				EventName:     e.EventName,
				Event:         e.Event,
				EventIndex:    e.EventIndex,
				Arguments:     arguments,
			})
			break loop
		skip:
		}
	}
	return
}

// 입력한 조건중 첫번째것을 반환
func (u *EventLogsUnpacker) FindOne(ca interface{}, eventID interface{}, filter Arguments) *UnpackedOutput {
	args := u.FindMany(ca, eventID, filter)
	if len(args) > 0 {
		return args[0]
	}
	return nil
}

// 입력을 common.Address로 변환
func toAddress(input interface{}) common.Address {
	switch v := input.(type) {
	case common.Address:
		return v
	case [20]byte:
		return common.BytesToAddress(v[:])
	case []byte:
		return common.BytesToAddress(v)
	case string:
		return common.HexToAddress(v)
	case nil:
		return common.Address{}
	default:
		panic(fmt.Errorf("not support type :%T as address", input))
	}
}

// 입력을 common.Hash로 변환
func toHash(input interface{}) common.Hash {
	switch v := input.(type) {
	case common.Hash:
		return v
	case [32]byte:
		return common.BytesToHash(v[:])
	case []byte:
		return common.BytesToHash(common.RightPadBytes(v, 32))
	case string:
		return common.HexToHash(v)
	case nil:
		return common.Hash{}
	default:
		panic(fmt.Errorf("not support type :%T as hash", input))
	}
}
