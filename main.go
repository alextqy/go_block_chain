package main

import (
	"fmt"
	b "go_block_chain/tools_basic"
	pow "go_block_chain/tools_pow"
	"time"
)

func main() {
	// basic
	blockchain1 := b.CreateBlockChain()
	time.Sleep(time.Second)
	blockchain1.AddBlock("After genesis, I have something to say.")
	time.Sleep(time.Second)
	blockchain1.AddBlock("Leo Cao is awesome!")
	time.Sleep(time.Second)
	blockchain1.AddBlock("I can't wait to follow his github!")
	time.Sleep(time.Second)
	for _, block := range blockchain1.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
	}

	// PoW
	blockchain2 := pow.CreateBlockChain()
	time.Sleep(time.Second)
	blockchain2.AddBlock("After genesis, I have something to say.")
	time.Sleep(time.Second)
	blockchain2.AddBlock("Leo Cao is awesome!")
	time.Sleep(time.Second)
	blockchain2.AddBlock("I can't wait to follow his github!")
	time.Sleep(time.Second)
	for _, block := range blockchain2.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
	}
}
