package block

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"time"
)
//	区块迭代器
type BlockchainIterator struct {
	CurrentHash []byte
	DB *bolt.DB	//
}

//	迭代器
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.Tip,bc.DB}
}


func (bi *BlockchainIterator) Next()*Block  {
	var block *Block
	err := bi.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil{
			currentBlockBytes := b.Get(bi.CurrentHash)
			//	获取当前迭代器里面的currentHash所对应的区块
			block = DeserializeBlick(currentBlockBytes)

			//	更新迭代器里面的CurrentHash
			bi.CurrentHash = block.PrevBlockHash
		}
		return nil

	})
	if err != nil{
		log.Panic(err)
	}
	return  block
}

//	遍历所有区块信息
func (blc *Blockchain) Printchain()  {
	blockchainIterator := blc.Iterator()

	for  {
		//	获取每个最新区块
		block := blockchainIterator.Next()

		fmt.Printf("Height: %d\n",block.Height)
		fmt.Printf("PrevBlockHash: %x\n",block.PrevBlockHash)
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Printf("Timestamp: %s\n",time.Unix(block.Timestamp,0).Format("2006-01-02 15:04:05"))
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Printf("Nonce: %d\n",block.Nonce)
		fmt.Println("--------------------------------------------------------------")

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		//	找到最后一个
		if big.NewInt(0).Cmp(&hashInt) == 0{
			break
		}
	}
}