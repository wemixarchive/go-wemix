package sim

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
- Full Log Search
sim.NewEventLogsUnpacker(receipt).FindMany(nil, nil, nil)
- Logs of a particular contract
sim.NewEventLogsUnpacker(receipt).FindMany(contract, nil, nil)
- Search for events in a particular contract
sim.NewEventLogsUnpacker(receipt).FindMany(contract, eventID, nil)
- Search for all specific events
sim.NewEventLogsUnpacker(receipt).FindMany(nil, eventID, nil)
- filter option
sim.NewEventLogsUnpacker(receipt).FindOne(nil, eventID, sim.Arguments{
			"from":  common.Address{},
			"to":    address,
			"value": amount,
		}))
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
	MaxUint256 = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)

	allEventsLock = sync.RWMutex{}
	allEvents     = allEventsType{}    // eventID => abi.Event
	allContracts  = allContractsType{} // CA => contract.Alias

	// When contact is generated, the address is matched with alias when EOA is generated, and when converted into a string, the address is replaced with alias.
	// When converting to string, zero address is set to 0x0 change to default
	replacer = []string{
		common.Address{}.Hex(), "0x00",
		MaxUint256.String(), "uint256_max",
		new(big.Int).Sub(MaxUint256, common.Big1).String(), "uint256_max-1",
	}
)

func collectEvents(events map[string]abi.Event) error {
	for _, e := range events {
		func() {
			allEventsLock.Lock()
			defer allEventsLock.Unlock()
			for _, exist := range allEvents[e.ID] {
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

type Arguments map[string]interface{}

type UnpackedOutput struct {
	ContractAlias string
	Contract      common.Address
	EventName     string
	Event         common.Hash
	EventIndex    uint
	Arguments     Arguments
}

func (u *UnpackedOutput) String() string {
	arguments, _ := json.MarshalIndent(u.Arguments, "", "  ")
	argumentsReplaced := strings.NewReplacer(replacer...).Replace(string(arguments))

	return fmt.Sprintf("%d [%s] %s%v", u.EventIndex, u.ContractAlias, u.EventName, argumentsReplaced)
}

type Unpacked struct {
	ContractAlias  string
	Contract       common.Address
	EventName      string
	Event          common.Hash
	EventIndex     uint
	ArgumentsSlice []Arguments
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
		var indexed, nonIndexed abi.Arguments

		for _, arg := range fn.Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			} else {
				nonIndexed = append(nonIndexed, arg)
			}
		}

		if len(e.Data) == 0 {
			if len(nonIndexed) > 0 {
				continue
			}
		} else {
			// unpack data
			if marshalledValues, err := fn.Inputs.UnpackValues(e.Data); err != nil {
				continue
			} else {
				for i, arg := range nonIndexed {
					arguments[arg.Name] = marshalledValues[i]
				}
			}
		}

		// unpack topic
		if len(e.Topics[1:]) != len(indexed) {
			continue
		} else if err := abi.ParseTopicsIntoMap(arguments, indexed, e.Topics[1:]); err != nil {
			continue
		}
		event = &fn
		// common.Hash
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

func (u *EventLogsUnpacker) unpack() {
	for _, e := range u.Receipt.Logs {
		argumentsSlice, event := u.unpackLog(e)
		if event == nil {
			continue
		}
		alias, ok := allContracts[e.Address]
		if !ok {
			alias = common.Bytes2Hex(e.Address[:])
		}

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

// Search and return all logs that meet the conditions
// If ca == nil, all contracts, default type common.Address
// If eventID == nil, all events id, default type common.Hash
// filter is a condition for log arguments
func (u *EventLogsUnpacker) FindMany(ca interface{}, eventID interface{}, filter Arguments) (events UnpackedArgumentsOutput) {
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
		if ca != (common.Address{}) && ca != e.Contract {
			continue
		}
		if eventID != (common.Hash{}) && eventID != e.Event {
			continue
		}
	loop:
		for _, arguments := range e.ArgumentsSlice {
			for k, v := range filter {
				if a := arguments[k]; a == nil {
					goto skip
				} else if !reflect.DeepEqual(a, v) {
					goto skip
				}
			}
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

func (u *EventLogsUnpacker) FindOne(ca interface{}, eventID interface{}, filter Arguments) *UnpackedOutput {
	args := u.FindMany(ca, eventID, filter)
	if len(args) > 0 {
		return args[0]
	}
	return nil
}

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
