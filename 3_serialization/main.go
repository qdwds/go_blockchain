package main

import (
	"fmt"
	block2 "serialization/block"
)

func main()  {

	//1_blockchain := block2.CreateBlockchainGenesisBlock()
	//1_blockchain.AddBlockToBlockchain(
	//	int64(len(1_blockchain.Blocks) + 1),
	//	1_blockchain.Blocks[len(1_blockchain.Blocks) - 1].Hash,
	//	"32",
	//	)
	//for _, v := range 1_blockchain.Blocks {
	//	fmt.Println(v.PrevBlockHash)
	//	fmt.Println(v.Hash)
	//}

	//block := block2.NewBlock(
	//	1,
	//	[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
	//	"32",
	//	)
	//fmt.Printf("%d\n",block.Nonce)
	//fmt.Printf("%v\n",block.Hash)
	//
	//2_pow := block2.NewProofOfWork(block)
	//fmt.Println(2_pow.IsVaild())
	//

	//bc := block.CreateBlockchainGenesisBlock("Genesis BLock...")
	//bc.AddBlockToBlockchain(
	//	int64(len(bc.Blocks) + 1),
	//	bc.Blocks[len(bc.Blocks) - 1].Hash,
	//	"32",
	//	)
	//for _, v := range bc.Blocks {
	//	fmt.Println(v.PrevBlockHash)
	//	fmt.Println(v.Hash)
	//	//2_pow := block.NewProofOfWork(v)
	//	//fmt.Printf("PoW: %s\n", strconv.FormatBool(2_pow.IsValid()))
	//	//fmt.Println(*2_pow)
	//}

	block := block2.NewBlock(1,make([] byte,32,32),"Ddaads")
	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)

	//	区块序列化
	bytes := block.Serialize()
	fmt.Println("序列化数据",bytes)

	//	反序列化
	block = block2.DeserializeBlick(bytes)
	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)
}
