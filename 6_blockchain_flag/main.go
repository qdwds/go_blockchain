package main

import "blockchain_bolt/block"

func main()  {
	blockchain := block.CreateBlockchainGenesisBlock("123")
	cli := block.Cli{blockchain}
	cli.Run()
}
