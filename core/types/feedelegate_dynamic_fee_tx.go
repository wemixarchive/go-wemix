// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type FeeDelegateDynamicFeeTx struct {
	SenderTx DynamicFeeTx
	FeePayer *common.Address `rlp:"nil"`
	// Signature values
	FV *big.Int `json:"fv" gencodec:"required"` // feePayer V
	FR *big.Int `json:"fr" gencodec:"required"` // feePayer R
	FS *big.Int `json:"fs" gencodec:"required"` // feePayer S
}

func (tx *FeeDelegateDynamicFeeTx) SetSenderTx(senderTx DynamicFeeTx) {
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
