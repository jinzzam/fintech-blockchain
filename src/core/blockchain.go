package core

import (
	"crypto/sha256"
	"time"
)

var GlobalBlockchains []*Blockchain

type Blockchain struct {
	Blocks           []Block
	BlockchainHeight uint64
	GenesisBlock     *Block
	CandidateBlock   *Block
}

func NewBlockchain() *Blockchain {
	b := newGenesisBlock()
	bc := &Blockchain{
		Blocks:           []Block{*b},
		BlockchainHeight: 1,
		GenesisBlock:     b,
		CandidateBlock:   nil,
	}
	return bc
}

func newGenesisBlock() *Block {
	b := &Block{
		Header: BlockHeader{
			PreviousHash: sha256.Sum256([]byte{}),
			Timestamp:    time.Now().UnixNano(),
			Index:        0,
		},
		Body: BlockBody{},
	}
	return b
}

func (bc *Blockchain) AddBlock() error {
	bc.Blocks = append(bc.Blocks, *bc.CandidateBlock)
	bc.BlockchainHeight = bc.BlockchainHeight + 1
	bc.CandidateBlock = nil

	return nil
}

func AppendBlockchain(bc *Blockchain) error {
	GlobalBlockchains = append(GlobalBlockchains, bc)
	return nil
}
