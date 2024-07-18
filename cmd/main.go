package main

import (
	"GoRestBlockchain/pkg/blockchain"
	"fmt"
	"strconv"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("This contains an image")

	bc.AddBlock("Pablo sent 1 BTC to the void")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		fmt.Println()
	}

}
