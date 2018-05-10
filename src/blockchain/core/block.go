package core

import (
	"encoding/binary"
	"time"
)

type BlockHeader struct {
	PreviousHash [32]byte
	Timestamp    int64
	Index        uint64
}

type BlockBody struct {
	Transactions []*Transaction
}

type Block struct {
	Header BlockHeader
	Body   BlockBody
}

func NewBlock(pb *Block) *Block {
	b := &Block{
		Header: BlockHeader{
			PreviousHash: Sha256.Sum256(pb.Header.ToBytes()),
			Timestamp:    time.Now().UnixNano(),
			Index:        pb.Header.Index + 1,
		},
		Body: BlockBody{
			Transactions: nil,
		},
	}
	return b
}

func (bh *BlockHeader) ToBytes() []byte {
	res := make([]byte, 0)
	tb := make([]byte, binary.MaxVarintLen64)
	ib := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(tb, bh.Timestamp)
	binary.PutUvarint(ib, bh.Index)
	for -,b := range [][]byte(bh.PreviousHash[:], tb, ib) {
		res = append(res, b...)
	}
  return res
}
