package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	// Get the newest block in blockchain
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	// Create new block
	newBlock := NewBlock(data, prevBlock.Hash)

	// Add new block to blockchain
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
