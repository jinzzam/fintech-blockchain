package core

import (
	"crypto/sha256"
	"encoding/binary"
	"time"

	"blockchain/validate"
)

type BlockHeader struct {
	PreviousHash   [32]byte
	MerkleRootHash [32]byte
	Timestamp      int64
	Index          uint64
}

func (bh *BlockHeader) ToBytes() []byte {
	res := make([]byte, 0)

	tb := make([]byte, binary.MaxVarintLen64)
	ib := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(tb, bh.Timestamp)
	binary.PutUvarint(ib, bh.Index)

	for _, b := range [][]byte{bh.PreviousHash[:], bh.MerkleRootHash[:], tb, ib} {
		res = append(res, b...)
	}

	return res
}

type BlockBody struct {
	Transactions []*Transaction
}

type Block struct {
	Header BlockHeader
	Body   BlockBody
}

func NewBlock(pb *Block) *Block {
	var diff [32]byte
	diff[0] = 10
	b := &Block{
		Header: BlockHeader{
			PreviousHash: sha256.Sum256(pb.Header.ToBytes()),
			Timestamp:    time.Now().UnixNano(),
			Index:        pb.Header.Index + 1,
		},
		Body: BlockBody{
			Transactions: nil,
		},
	}
	return b
}

func (b *Block) AddTransaction(t *Transaction) error {
	b.Body.Transactions = append(b.Body.Transactions, t)
	tb := make([][]byte, 0)
	for _, v := range b.Body.Transactions {
		tb = append(tb, v.ToBytes())
	}
	b.Header.MerkleRootHash = validate.MerkleRootHash(tb)
	return nil
}
