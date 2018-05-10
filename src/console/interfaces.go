package console

// Blockchainer is a interface to perform basic operations for a block chain.
type Blockchainer interface {
	AddBlock() error
	String() string
}

// Blocker is a interface of Block.
type Blocker interface {
	Block() interface{}
	String() string
}
