package core

import (
	"crypto/sha256"
	"time"
)

var GlobalBlockChains []*Blockchain

type Blockchain struct { //블록체인 구조
	Blocks           []Block
	BlockChainHeight uint64
	GenesisBlock     *Block
	CandidateBlock   *Block
}

func NewBlockChain() *Blockchain { //블록체인 생성
	b := newGenesisBlock()
	bc := &Blockchain{
		Blocks:           []Block{*b},
		BlockChainHeight: 1,
		GenesisBlock:     b,
		CandidateBlock:   nil,
	}
	return bc
}

func newGenesisBlock() *Block { //제네시스 블록 생성
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

func (bc *Blockchain) AddBlock() error { //후보 블록 생성
	bc.Blocks = append(bc.Blocks, *bc.CandidateBlock)
	bc.BlockChainHeight = bc.BlockChainHeight + 1
	bc.CandidateBlock = nil
	return nil
}

func AppendBlockchain(bc *Blockchain) { //블록체인 발견 시 추가
	GlobalBlockChain = append(GlobalBlockChain, bc)
	return nil
}
