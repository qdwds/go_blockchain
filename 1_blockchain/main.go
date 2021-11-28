package main

import (
	"fmt"
	block2 "new_block/block"
)

func main()  {

	blockchain := block2.CreateBlockchainGenesisBlock()
	blockchain.AddBlockToBlockchain(
		int64(len(blockchain.Blocks) + 1),
		blockchain.Blocks[len(blockchain.Blocks) - 1].Hash,
		"32",
		)
	for _, v := range blockchain.Blocks {
		fmt.Println(v.PrevBlockHash)
		fmt.Println(v.Hash)
	}
}
