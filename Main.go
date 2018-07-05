package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("Block 1 ne")
	bc.AddBlock("Block 2 ne")
	bc.AddBlock("Block 3 ne")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Printf("Current hash: %x\n", block.Hash)
		fmt.Printf("Data real: %s\n", block.Data)
		fmt.Println()
	}
}
