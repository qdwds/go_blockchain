package main

import (
	block2 "bolt/block"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
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
	//
	////	区块序列化
	//bytes := block.Serialize()
	//fmt.Println("序列化数据",bytes)
	//
	////	反序列化
	//block = block2.DeserializeBlick(bytes)
	//fmt.Printf("%d\n",block.Nonce)
	//fmt.Printf("%x\n",block.Hash)

	//	创建数据库
	db,err := bolt.Open("my.db",0600,nil)
	defer db.Close()
	if err != nil{
		log.Fatal(err)
	}
	//db.Update(func(tx *bolt.Tx) error {
	//	//	取表对象
	//	b :=tx.Bucket([]byte("blocks"))
	//	//	如果表不存在就创建表
	//	if b == nil{
	//		b,err = tx.CreateBucket([]byte("block"))
	//		if err != nil {
	//			log.Panic("Block table create  failed")
	//		}
	//	}
	//	//	往表中插入`反序列化的区块数据`
	//	err := b.Put([]byte("l"),block.Serialize())
	//	if err != nil{
	//		log.Panic(err)
	//	}
	//	return nil
	//})
	//if err != nil {
	//	log.Panic(err )
	//}


	//	查看
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("block"))
		if b != nil {
			blockData := b.Get([]byte("l"))
			block := block2.DeserializeBlick(blockData)
			fmt.Printf("%v\n",block)
		}
		return nil
	})
	if err != nil {
		log.Panic(err )
	}
}
