package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func calculateHash(b *Block) [32]byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	return hash
}

func (b *Block) setHash() {
	hash := calculateHash(b)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	newBlock := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
		0,
	}

	pow := NewProofOfWork(newBlock)

	nonce, hash := pow.Run()

	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce

	//newBlock.setHash()

	return newBlock

}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
