package ethtest

import "github.com/ethereum/go-ethereum/eth/protocols/snap"

// GetAccountRange represents an account range query.
type GetAccountRange snap.GetAccountRangePacket

func (g GetAccountRange) Code() int { return 39 }

type AccountRange snap.AccountRangePacket

func (g AccountRange) Code() int { return 40 }

type GetStorageRanges snap.GetStorageRangesPacket

func (g GetStorageRanges) Code() int { return 41 }

type StorageRanges snap.StorageRangesPacket

func (g StorageRanges) Code() int { return 42 }

type GetByteCodes snap.GetByteCodesPacket

func (g GetByteCodes) Code() int { return 43 }

type ByteCodes snap.ByteCodesPacket

func (g ByteCodes) Code() int { return 44 }

type GetTrieNodes snap.GetTrieNodesPacket

func (g GetTrieNodes) Code() int { return 45 }

type TrieNodes snap.TrieNodesPacket

func (g TrieNodes) Code() int { return 46 }
