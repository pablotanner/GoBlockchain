package blockchain

import (
	"fmt"
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	Tip []byte
	Db  *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		// Get the blocks Bucket
		b := tx.Bucket([]byte(blocksBucket))
		// Retrieve hash of last block (Tip)
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		fmt.Printf("Failed to get last Hash from DB")
	}
	newBlock := NewBlock(data, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		// Get the blocks Bucket
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			fmt.Printf("Failed to add block hash to chain")
		}
		// Update "l" key, which stores most recent block's hash
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			fmt.Printf("Failed to update l key in DB")
		}
		bc.Tip = newBlock.Hash

		return nil
	})

}

func NewBlockchain() *Blockchain {
	// Hash of the last block
	var tip []byte
	// Open BoltDB File
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		fmt.Printf("Failed to open BoltDB File")
	}

	// Open read-write transaction
	err = db.Update(func(tx *bolt.Tx) error {
		// Try to retrieve the bucket "blocksBucket"
		b := tx.Bucket([]byte(blocksBucket))

		// If it couldn't find it, create it and make it store a new genesis block
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))

			if err != nil {
				fmt.Printf("Failed to create Bucket")
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			// Otherwise we get the hash of the last block in the bucket
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	bc := Blockchain{tip, db}

	return &bc
}
