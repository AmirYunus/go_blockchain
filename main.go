package main

import (
	"fmt"
	"github.com/AmirYunus/go_blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First")
	chain.AddBlock("Second")
	chain.AddBlock("Third")

	for index, block := range chain.Blocks {
		fmt.Printf("Block: %x\n", index)
		// fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
