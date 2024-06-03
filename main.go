package main

import (
	"fmt"
	"strconv"

	"github.com/pedro77h/blockchain-go/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlock("New Transaction of $500!")
	chain.AddBlock("New Transaction of $600!")
	chain.AddBlock("New Transaction of $700!")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash in block: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
