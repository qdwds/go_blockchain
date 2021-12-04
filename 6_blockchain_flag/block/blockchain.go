package block

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

//	数据库名字
const dbName = "blockchaom.db"
//	数据库 表名
const blockTableName = "blocks"

//	区块链结构体
type Blockchain struct {
	Tip []byte 	//	最新区块的hash值
	DB *bolt.DB	//
}

//	判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(dbName);os.IsNotExist(err) {
		return false
	}
	return  true
}

//	创建带有传世区块的区块链
func CreateBlockchainGenesisBlock(data string) *Blockchain {
	fmt.Println(dbExists())
	//	校样传世区块是否存在
	if dbExists(){
		fmt.Println("创建区块已经存在")
		db,err := bolt.Open(dbName,0600,nil)
		if err != nil{
			log.Panic(err,77777)
		}
		var blockchain *Blockchain

		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blockTableName))
			hash := b.Get([]byte("l"))
			blockchain = &Blockchain{hash,db}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		return blockchain
	}

	//	创建传世区块
	db,err := bolt.Open(dbName,0600,nil)
	if err != nil{
		log.Panic(err)
	}
	var blockHash []byte
	err = db.Update(func(tx *bolt.Tx) error {
		//	首先校样表是否存在
		b := tx.Bucket([]byte(blockTableName))
		if b == nil{
			//	创世区块中直接创建
			b,err = tx.CreateBucket([]byte(blockTableName))
			if err != nil{
				log.Panic(err)
			}
		}


		//	创建创世区块
		if b != nil{
			gb := CreateGenesisBlock(data)
			//	hash 作为key值
			err := b.Put(gb.Hash,gb.Serialize())
			if err != nil{
				log.Panic(err)
			}

			//	存储最新的区块的hash
			err = b.Put([]byte("l"),gb.Hash)
			if err != nil{
				log.Panic(err)
			}
			blockHash = gb.Hash
		}

		return nil
	})

	//	创建传世区块
	//genesisBlock := CreateGenesisBlock(data)
	//	返回区块对象
	return &Blockchain{blockHash,db}
}



//	添加区块到区块链中
func (blc *Blockchain) AddBlockToBlockchain( data string)  {

	////	往数区块链数组中添加新区块
	//bc.Blocks = append(bc.Blocks, newBlock)
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//	获取表
		b := tx.Bucket([]byte(blockTableName))
		//	创建新区块
		if b != nil {
			//	获取当前最新的区块信息
			blockBytes := b.Get(blc.Tip)
			//	反序列化 获取最新区块数据
			block := DeserializeBlick(blockBytes)
			//	将区块序列化 =>  存储到数据库中
			newBlock := NewBlock(block.Height + 1, block.Hash, data)
			err := b.Put(newBlock.Hash,newBlock.Serialize())
			if err != nil{
				log.Panic(err)
			}
			//	更新数据库中"l"对应的hash
			err = b.Put([]byte("l"),newBlock.Hash)
			if err != nil{
				log.Panic(err)
			}

			//	更新blockchain到Tip
			blc.Tip = newBlock.Hash
		}

		return nil
	})
	if err != nil{
		log.Panic(err)
	}
}
