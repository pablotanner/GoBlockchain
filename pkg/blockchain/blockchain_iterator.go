package blockchain

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type ChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterator() *ChainIterator {
	ci := &ChainIterator{bc.Tip, bc.Db}

	return ci
}

func (i *ChainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		fmt.Printf("Failed to read DB for ChainIterator")
	}

	i.currentHash = block.PrevBlockHash

	return block
}
