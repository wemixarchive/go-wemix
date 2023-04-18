## Implementation of Fee Delegation(FD) Transaction Structure

Fee Delegation Transaction is a service where the feePayer pays the fee for the transaction that the Sender wants to execute. This is done by adding the feePayer's signature to the existing transaction information signed by the Sender and sending it.
Fee Delegation Transaction only supports the DynamicFeeTxType among the existing transactions signed by the Sender and does not support LegacyTxType or AccessListTxType.


|      Tx Type       | Tx Type supported with fee delegation |
|:------------------:|---------------------------------------|
|    LegacyTxType    | No                                    |
|  AccessListTxType  | No                                    |
|  DynamicFeeTxType  | Yes                                   |


***
### 1.Add FD Transaction Types

* core/types/transaction.go

``` go
const (
	LegacyTxType = iota
	AccessListTxType
	DynamicFeeTxType
	FeeDelegateDynamicFeeTxType = 22 //fee delegation
)
```

### 2.Add feePayer in Transaction struct

* core/types/transaction.go

``` go
type Transaction struct {
	inner TxData    // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash     atomic.Value
	size     atomic.Value
	from     atomic.Value
	feePayer atomic.Value  //fee delegation
}
```

### 3.Add FD interface in TxData

* core/types/transaction.go

``` go
type TxData interface {
	txType() byte // returns the type ID
	copy() TxData // creates a deep copy and initializes all fields

	chainID() *big.Int
	accessList() AccessList
	data() []byte
	gas() uint64
	gasPrice() *big.Int
	gasTipCap() *big.Int
	gasFeeCap() *big.Int
	value() *big.Int
	nonce() uint64
	to() *common.Address
	rawSignatureValues() (v, r, s *big.Int)
	setSignatureValues(chainID, v, r, s *big.Int)
	// fee delegation
	feePayer() *common.Address
	rawFeePayerSignatureValues() (v, r, s *big.Int)
}
```

### 5.Add New FD Transaction Structure

* core/types/feedelegate_dynamic_fee_tx.go

``` go
type FeeDelegateDynamicFeeTx struct {
	SenderTx DynamicFeeTx
	FeePayer *common.Address `rlp:"nil"`
	// Signature values
	FV *big.Int `json:"fv" gencodec:"required"` // feePayer V
	FR *big.Int `json:"fr" gencodec:"required"` // feePayer R
	FS *big.Int `json:"fs" gencodec:"required"` // feePayer S
}
```

### 6.Add FeeDelegateDynamicFeeTx getter and setter

* core/types/feedelegate_dynamic_fee_tx.go

``` go
func (tx *FeeDelegateDynamicFeeTx) SetSenderTx(senderTx TxData) {
	tx.SenderTx.ChainID = senderTx.ChainID
	tx.SenderTx.Nonce = senderTx.Nonce
	tx.SenderTx.GasFeeCap = senderTx.GasFeeCap
	tx.SenderTx.GasTipCap = senderTx.GasTipCap
	tx.SenderTx.Gas = senderTx.Gas
	tx.SenderTx.To = senderTx.To
	tx.SenderTx.Value = senderTx.Value
	tx.SenderTx.Data = senderTx.Data
	copy(tx.SenderTx.AccessList, senderTx.AccessList)

	v, r, s := senderTx.rawSignatureValues()
	tx.SenderTx.V = v
	tx.SenderTx.R = r
	tx.SenderTx.S = s
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *FeeDelegateDynamicFeeTx) copy() TxData {
	cpy := &FeeDelegateDynamicFeeTx{
		SenderTx: tx.copySenderTx(),
		FeePayer: copyAddressPtr(tx.FeePayer),
		FV:       new(big.Int),
		FR:       new(big.Int),
		FS:       new(big.Int),
	}
	if tx.FV != nil {
		cpy.FV.Set(tx.FV)
	}
	if tx.FR != nil {
		cpy.FR.Set(tx.FR)
	}
	if tx.FS != nil {
		cpy.FS.Set(tx.FS)
	}
	return cpy
}

func (tx *FeeDelegateDynamicFeeTx) copySenderTx() DynamicFeeTx {
	cpy := DynamicFeeTx{
		Nonce: tx.SenderTx.Nonce,
		To:    copyAddressPtr(tx.SenderTx.To),
		Data:  common.CopyBytes(tx.SenderTx.Data),
		Gas:   tx.SenderTx.Gas,
		// These are copied below.
		AccessList: make(AccessList, len(tx.SenderTx.AccessList)),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
		GasTipCap:  new(big.Int),
		GasFeeCap:  new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	copy(cpy.AccessList, tx.SenderTx.accessList())
	if tx.SenderTx.Value != nil {
		cpy.Value.Set(tx.SenderTx.Value)
	}
	if tx.SenderTx.ChainID != nil {
		cpy.ChainID.Set(tx.SenderTx.ChainID)
	}
	if tx.SenderTx.GasTipCap != nil {
		cpy.GasTipCap.Set(tx.SenderTx.GasTipCap)
	}
	if tx.SenderTx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.SenderTx.GasFeeCap)
	}
	if tx.SenderTx.V != nil {
		cpy.V.Set(tx.SenderTx.V)
	}
	if tx.SenderTx.R != nil {
		cpy.R.Set(tx.SenderTx.R)
	}
	if tx.SenderTx.S != nil {
		cpy.S.Set(tx.SenderTx.S)
	}
	return cpy
}

// accessors for innerTx.
func (tx *FeeDelegateDynamicFeeTx) txType() byte              { return FeeDelegateDynamicFeeTxType }
func (tx *FeeDelegateDynamicFeeTx) chainID() *big.Int         { return tx.SenderTx.ChainID }
func (tx *FeeDelegateDynamicFeeTx) accessList() AccessList    { return tx.SenderTx.AccessList }
func (tx *FeeDelegateDynamicFeeTx) data() []byte              { return tx.SenderTx.Data }
func (tx *FeeDelegateDynamicFeeTx) gas() uint64               { return tx.SenderTx.Gas }
func (tx *FeeDelegateDynamicFeeTx) gasFeeCap() *big.Int       { return tx.SenderTx.GasFeeCap }
func (tx *FeeDelegateDynamicFeeTx) gasTipCap() *big.Int       { return tx.SenderTx.GasTipCap }
func (tx *FeeDelegateDynamicFeeTx) gasPrice() *big.Int        { return tx.SenderTx.GasFeeCap }
func (tx *FeeDelegateDynamicFeeTx) value() *big.Int           { return tx.SenderTx.Value }
func (tx *FeeDelegateDynamicFeeTx) nonce() uint64             { return tx.SenderTx.Nonce }
func (tx *FeeDelegateDynamicFeeTx) to() *common.Address       { return tx.SenderTx.To }
func (tx *FeeDelegateDynamicFeeTx) feePayer() *common.Address { return tx.FeePayer }
func (tx *FeeDelegateDynamicFeeTx) rawFeePayerSignatureValues() (v, r, s *big.Int) {
	return tx.FV, tx.FR, tx.FS
}

func (tx *FeeDelegateDynamicFeeTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.SenderTx.rawSignatureValues()
}

func (tx *FeeDelegateDynamicFeeTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.FV, tx.FR, tx.FS = v, r, s
}
```

### 7.Add Fee Delegation Signier

* Include core/types/transaction_signing.go

``` go
// fee delegation
type feeDelegateSigner struct{ londonSigner }

func NewFeeDelegateSigner(chainId *big.Int) Signer {
	return feeDelegateSigner{londonSigner{eip2930Signer{NewEIP155Signer(chainId)}}}
}

func (s feeDelegateSigner) Sender(tx *Transaction) (common.Address, error) {
	if tx.Type() != FeeDelegateDynamicFeeTxType {
		return s.londonSigner.Sender(tx)
	}
	V, R, S := tx.RawFeePayerSignatureValues()
	// DynamicFee txs are defined to use 0 and 1 as their recovery
	// id, add 27 to become equivalent to unprotected Homestead signatures.
	V = new(big.Int).Add(V, big.NewInt(27))
	if tx.ChainId().Cmp(s.chainId) != 0 {
		return common.Address{}, ErrInvalidChainId
	}
	return recoverPlain(s.Hash(tx), R, S, V, true)
}

func (s feeDelegateSigner) Equal(s2 Signer) bool {
	x, ok := s2.(feeDelegateSigner)
	return ok && x.chainId.Cmp(s.chainId) == 0
}

func (s feeDelegateSigner) SignatureValues(tx *Transaction, sig []byte) (R, S, V *big.Int, err error) {
	txdata, ok := tx.inner.(*FeeDelegateDynamicFeeTx)

	if !ok {
		return s.londonSigner.SignatureValues(tx, sig)
	}
	// Check that chain ID of tx matches the signer. We also accept ID zero here,
	// because it indicates that the chain ID was not specified in the tx.
	if txdata.SenderTx.chainID().Sign() != 0 && txdata.SenderTx.chainID().Cmp(s.chainId) != 0 {
		return nil, nil, nil, ErrInvalidChainId
	}
	R, S, _ = decodeSignature(sig)
	V = big.NewInt(int64(sig[64]))
	return R, S, V, nil
}

// Hash returns the hash to be signed by the sender.
// It does not uniquely identify the transaction.
func (s feeDelegateSigner) Hash(tx *Transaction) common.Hash {
	senderV, senderR, senderS := tx.RawSignatureValues()
	return prefixedRlpHash(
		tx.Type(),
		[]interface{}{
			[]interface{}{
				s.chainId,
				tx.Nonce(),
				tx.GasTipCap(),
				tx.GasFeeCap(),
				tx.Gas(),
				tx.To(),
				tx.Value(),
				tx.Data(),
				tx.AccessList(),
				senderV,
				senderR,
				senderS,
			},
			tx.FeePayer(),
		})
}
```
***
***
## Extend and Add RPC api

### 1. Add feePayer,FV,FR,FS in RPCTransaction Struct 
* internal/ethapi/api.go

``` go
type RPCTransaction struct {
	BlockHash        *common.Hash      `json:"blockHash"`
	BlockNumber      *hexutil.Big      `json:"blockNumber"`
	From             common.Address    `json:"from"`
	Gas              hexutil.Uint64    `json:"gas"`
	GasPrice         *hexutil.Big      `json:"gasPrice"`
	GasFeeCap        *hexutil.Big      `json:"maxFeePerGas,omitempty"`
	GasTipCap        *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
	Hash             common.Hash       `json:"hash"`
	Input            hexutil.Bytes     `json:"input"`
	Nonce            hexutil.Uint64    `json:"nonce"`
	To               *common.Address   `json:"to"`
	TransactionIndex *hexutil.Uint64   `json:"transactionIndex"`
	Value            *hexutil.Big      `json:"value"`
	Type             hexutil.Uint64    `json:"type"`
	Accesses         *types.AccessList `json:"accessList,omitempty"`
	ChainID          *hexutil.Big      `json:"chainId,omitempty"`
	V                *hexutil.Big      `json:"v"`
	R                *hexutil.Big      `json:"r"`
	S                *hexutil.Big      `json:"s"`
	// fee delegation
	FeePayer *common.Address `json:"feePayer,omitempty"`
	FV       *hexutil.Big    `json:"fv,omitempty"`
	FR       *hexutil.Big    `json:"fr,omitempty"`
	FS       *hexutil.Big    `json:"fs,omitempty"`
}
```

### 2.Extend signTransaction RPC

- example :

```
$ bin/gweimx.sh console
> personal.signTransaction({
 from:"0x82667998ae5fd9e4f4637fc805e97740c673c517",
 chainId: "0x458",
 gas: "0x5208",
 maxFeePerGas: "0x174876e801",
 maxPriorityFeePerGas: "0x174876e801",
 hash: "0x19b5e90deb03cd4b87aca41ca09dcfed5e7c7ad33e6579f30f7efba722c54424",
 input: "0x",    
 nonce: "0x1c",
 to: "0xdb8408bb47bf5e745fed00fc2c99e2f4e1a9270f",
 value: "0xdde0b6b3a7640000",
 r: "0xb4dc96b4580bd1d3f090b953ea3612625dd834af8e7cc6146def84a0c137b32c",
 s: "0x82427bf3a4589ffe79bebad9151ece587790462cd44abdd810abe0016134b8f",
 v: "0x0",
 feePayer: "0xdb8408bb47bf5e745fed00fc2c99e2f4e1a9270f"
},"test")
```
result:
```
{
 raw: "0x16f8d0f8768204581c85174876e80185174876e80182520894db8408bb47bf5e745fed00fc2c99e2f4e1a9270f880de0b6b3a764000080c080a0b4dc96b4580bd1d3f090b953ea3612625dd834af8e7cc6146def84a0c137b32ca0082427bf3a4589ffe79bebad9151ece587790462cd44abdd810abe0016134b8f94db8408bb47bf5e745fed00fc2c99e2f4e1a9270f80a0ba376dff9a2a4344a570367c94eeee2434a0e44ccf2da900f54becc6adaf0b5ca077b7d15ab7ba7213b5189094ca1bd41c7a48390767e4acb14c98f3442e561abb",
 tx: {
 accessList: [],
 chainId: "0x458",
 feePayer: "0xdb8408bb47bf5e745fed00fc2c99e2f4e1a9270f",
 fr: "0xba376dff9a2a4344a570367c94eeee2434a0e44ccf2da900f54becc6adaf0b5c",
 fs: "0x77b7d15ab7ba7213b5189094ca1bd41c7a48390767e4acb14c98f3442e561abb",
 fv: "0x0",
 gas: "0x5208",
 gasPrice: null,
 hash: "0x19b5e90deb03cd4b87aca41ca09dcfed5e7c7ad33e6579f30f7efba722c54424",
 input: "0x",
 maxFeePerGas: "0x174876e801",
 maxPriorityFeePerGas: "0x174876e801",
 nonce: "0x1c",
 r: "0xb4dc96b4580bd1d3f090b953ea3612625dd834af8e7cc6146def84a0c137b32c",
 s: "0x82427bf3a4589ffe79bebad9151ece587790462cd44abdd810abe0016134b8f",
 to: "0xdb8408bb47bf5e745fed00fc2c99e2f4e1a9270f",
 type: "0x16",
 v: "0x0",
 value: "0xde0b6b3a7640000"
 }
}
```
### 3.Add SignRawFeeDelegateTransaction RPC
* internal/ethapi/api.go

``` go
func (s *PrivateAccountAPI) SignRawFeeDelegateTransaction(ctx context.Context, args TransactionArgs, input hexutil.Bytes, passwd string) (*SignTransactionResult, error) {
	if args.FeePayer == nil {
		return nil, fmt.Errorf("missing FeePayer")
	}

	// Look up the wallet containing the requested signer
	account := accounts.Account{Address: *args.FeePayer}

	wallet, err := s.am.Find(account)
	if err != nil {
		return nil, err
	}

	rawTx := new(types.Transaction)
	if err := rawTx.UnmarshalBinary(input); err != nil {
		return nil, err
	}

	V, R, S := rawTx.RawSignatureValues()
	if rawTx.Type() == types.DynamicFeeTxType {
		SenderTx := types.DynamicFeeTx{
			To:         rawTx.To(),
			ChainID:    rawTx.ChainId(),
			Nonce:      rawTx.Nonce(),
			Gas:        rawTx.Gas(),
			GasFeeCap:  rawTx.GasFeeCap(),
			GasTipCap:  rawTx.GasTipCap(),
			Value:      rawTx.Value(),
			Data:       rawTx.Data(),
			AccessList: rawTx.AccessList(),
			V:          V,
			R:          R,
			S:          S,
		}

		FeeDelegateDynamicFeeTx := &types.FeeDelegateDynamicFeeTx{
			FeePayer: args.FeePayer,
		}

		FeeDelegateDynamicFeeTx.SetSenderTx(SenderTx)
		tx := types.NewTx(FeeDelegateDynamicFeeTx)

		signed, err := wallet.SignTxWithPassphrase(account, passwd, tx, s.b.ChainConfig().ChainID)

		if err != nil {
			return nil, err
		}
		data, err := signed.MarshalBinary()
		if err != nil {
			return nil, err
		}

		return &SignTransactionResult{data, signed}, nil
	}

	return nil, fmt.Errorf("senderTx type error")
}
```
- example :
```
$ bin/gweimx.sh console
> personal.signRawFeeDelegateTransaction({"feePayer":"0xdb8408bb47bf5e745fed00fc2c99e2f4e1a9270f"},
"0x02f8768204581c85174876e80185174876e80182520894db8408bb47bf5e745fed00fc2c99e2f4e1a9270f880de0b6b3a764000080c080a0b4dc96b4580bd1d3f090b953ea3612625dd834af8e7cc6146def84a0c137b32ca0082427bf3a4589ffe79bebad9151ece587790462cd44abdd810abe0016134b8f"
,"test")
```
result:
```
{
raw: "0x16f8d0f8768204581c85174876e80185174876e80182520894db8408bb47bf5e745fed00fc2c99e2f4e1a9270f880de0b6b3a764000080c080a0b4dc96b4580bd1d3f090b953ea3612625dd834af8e7cc6146def84a0c137b32ca0082427bf3a4589ffe79bebad9151ece587790462cd44abdd810abe0016134b8f94db8408bb47bf5e745fed00fc2c99e2f4e1a9270f80a0ba376dff9a2a4344a570367c94eeee2434a0e44ccf2da900f54becc6adaf0b5ca077b7d15ab7ba7213b5189094ca1bd41c7a48390767e4acb14c98f3442e561abb",
tx: {
 accessList: [],
 chainId: "0x458",
 feePayer: "0xdb8408bb47bf5e745fed00fc2c99e2f4e1a9270f",
 fr: "0xba376dff9a2a4344a570367c94eeee2434a0e44ccf2da900f54becc6adaf0b5c",
 fs: "0x77b7d15ab7ba7213b5189094ca1bd41c7a48390767e4acb14c98f3442e561abb",
 fv: "0x0",
 gas: "0x5208",
 gasPrice: null,
 hash: "0x19b5e90deb03cd4b87aca41ca09dcfed5e7c7ad33e6579f30f7efba722c54424",
 input: "0x",
 maxFeePerGas: "0x174876e801",
 maxPriorityFeePerGas: "0x174876e801",
 nonce: "0x1c",
 r: "0xb4dc96b4580bd1d3f090b953ea3612625dd834af8e7cc6146def84a0c137b32c",
 s: "0x82427bf3a4589ffe79bebad9151ece587790462cd44abdd810abe0016134b8f",
 to: "0xdb8408bb47bf5e745fed00fc2c99e2f4e1a9270f",
 type: "0x16",
 v: "0x0",
 value: "0xde0b6b3a7640000"
 }
}
```

### 4.Add feePayer information to getTransaction RPC result

``` 
$ bin/gwemix.sh console
> eth.getTransaction('0x53ef850a55c09b698ccc33bb5f373ea1e67f5002b4d48bdcf9b2078b368d36a8')
```
result:
```
{
  accessList: [],
  blockHash: "0x4cab269319725d2cf4d6eb118b11a54c6d73f819d01cc9c64653a5d1975254b8",
  blockNumber: 127551,
  chainId: "0x458",
  feePayer: "0x82667998ae5fd9e4f4637fc805e97740c673c517",
  fr: "0xde281f30f9c06f17a77ce391ef1f91b2c5bbcc28dbd4350c9a31c6d0b0fe5a37",
  from: "0xe6205771b7777421bfddb90c93eda3f3d9d6e35a",
  fs: "0x1f678163b0dfbd8913473305c9b3ee0e3d19705440df93d5095b1231f600bb61",
  fv: "0x1",
  gas: 21000,
  gasPrice: 100000000001,
  hash: "0x53ef850a55c09b698ccc33bb5f373ea1e67f5002b4d48bdcf9b2078b368d36a8",
  input: "0x",
  maxFeePerGas: 100000000001,
  maxPriorityFeePerGas: 100000000001,
  nonce: 30,
  r: "0xbf9e47ffaf39aed6523c93f94badd5df32618bfd041fd95f5c411cbe79b5f7f5",
  s: "0x5599c2eea6d4b290259ff545b65ba680cf4d1591610e04dca5b31b879acf09b",
  to: "0x4231d51941be79312203e4c029a3b56162902353",
  transactionIndex: 0,
  type: "0x16",
  v: "0x0",
  value: 1000000000000000000
}
```
### Fee Delegation Transaction Example
-  Reference code : https://github.com/wemixarchive/feedelegation-js