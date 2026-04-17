package main

import (
	"fmt"
)

func main() {
	genesis := NewGenesisBlock()
	fmt.Printf("--- Genesis Block ---\n")
	fmt.Printf("Hash: %x\n", genesis.Hash)
	fmt.Printf("Data: %s\n\n", genesis.Data)

	secondBlock := NewBlock("Identity: ABC_User (Person)", genesis.Hash)
	
	fmt.Printf("--- Second Block ---\n")
	fmt.Printf("Hash: %x\n", secondBlock.Hash)
	fmt.Printf("Prev Hash: %x\n", secondBlock.PrevBlockHash)
	fmt.Printf("Data: %s\n\n", secondBlock.Data)

	serialized := secondBlock.Serialize()
	fmt.Printf("Serialized bytes (Ready for Database): %d bytes\n", len(serialized))

	reconstructed := DeserializeBlock(serialized)
	fmt.Printf("Reconstructed Data: %s\n", reconstructed.Data)
}
