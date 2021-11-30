package block

import (
	"github.com/boltdb/bolt"
	"log"
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


//	创建带有传世区块的区块链
func CreateBlockchainGenesisBlock(data string) *Blockchain {
	db,err := bolt.Open(dbName,0600,nil)
	if err != nil{
		log.Panic(err)
	}
	var blckHash []byte
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
			blckHash = gb.Hash
		}

		return nil
	})

	//	创建传世区块
	//genesisBlock := CreateGenesisBlock(data)
	//	返回区块对象
	return &Blockchain{blckHash,db}
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

//	遍历 循环 获取所有区块信息
//func (blc *Blockchain)Printchain()  {
//	var block *Block
//	var currentHash []byte = blc.Tip
//	for {
//		err := blc.DB.View(func(tx *bolt.Tx) error {
//			b := tx.Bucket([]byte(blockTableName))
//			if b != nil {
//				//	获取当前区块的字节数组
//				blockBytes := b.Get([]byte(currentHash))
//				//	反序列化
//				block = DeserializeBlick(blockBytes)
//				fmt.Printf("Height: %d\n",block.Height)
//				fmt.Printf("PrevBlockHash: %x\n",block.PrevBlockHash)
//				fmt.Printf("Data: %s\n",block.Data)
//				fmt.Printf("Timestamp: %s\n",time.Unix(block.Timestamp,0).Format("2006-01-02 15:04:05"))
//				fmt.Printf("Hash: %x\n",block.Hash)
//				fmt.Printf("Nonce: %d\n",block.Nonce)
//				fmt.Println("--------------------------------------------------------------")
//			}
//
//			return nil
//		})
//		if err != nil{
//			log.Panic(err)
//		}
//
//		var hashInt big.Int
//		hashInt.SetBytes(block.PrevBlockHash)
//		if big.NewInt(0).Cmp(&hashInt) == 0{
//			break
//		}
//		currentHash = block.PrevBlockHash
//		currentHash = block.PrevBlockHash
//	}
//}