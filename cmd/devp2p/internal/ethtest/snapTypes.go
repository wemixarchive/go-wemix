// Copyright 2022 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package ethtest

import "github.com/ethereum/go-ethereum/eth/protocols/snap"

// GetAccountRange represents an account range query.
type GetAccountRange snap.GetAccountRangePacket

func (g GetAccountRange) Code() int       { return 39 }
func (msg GetAccountRange) ReqID() uint64 { return msg.ID }

type AccountRange snap.AccountRangePacket

func (g AccountRange) Code() int       { return 40 }
func (msg AccountRange) ReqID() uint64 { return msg.ID }

type GetStorageRanges snap.GetStorageRangesPacket

func (g GetStorageRanges) Code() int       { return 41 }
func (msg GetStorageRanges) ReqID() uint64 { return msg.ID }

type StorageRanges snap.StorageRangesPacket

func (g StorageRanges) Code() int       { return 42 }
func (msg StorageRanges) ReqID() uint64 { return msg.ID }

type GetByteCodes snap.GetByteCodesPacket

func (g GetByteCodes) Code() int       { return 43 }
func (msg GetByteCodes) ReqID() uint64 { return msg.ID }

type ByteCodes snap.ByteCodesPacket

func (g ByteCodes) Code() int       { return 44 }
func (msg ByteCodes) ReqID() uint64 { return msg.ID }

type GetTrieNodes snap.GetTrieNodesPacket

func (g GetTrieNodes) Code() int       { return 45 }
func (msg GetTrieNodes) ReqID() uint64 { return msg.ID }

type TrieNodes snap.TrieNodesPacket

func (g TrieNodes) Code() int       { return 46 }
func (msg TrieNodes) ReqID() uint64 { return msg.ID }
