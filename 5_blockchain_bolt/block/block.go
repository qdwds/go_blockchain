package block

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

//	区块结构体
type Block struct {
	Height int64			//	区块高度
	PrevBlockHash []byte	//	上一个区块的HASH
	Data []byte				//	交易数据
	Timestamp int64			//	时间戳
	Hash []byte				//	当前HASH
	Nonce int64				//	工作量证明
}



//	创建新 区块
func NewBlock( height int64, PrevBlockHash []byte,data string) *Block {
  	block := &Block{height, PrevBlockHash, []byte(data),time.Now().Unix(), nil,0}
	//	调用工作量证明的方法 返回有效的Hash和Nonce
	pow := NewProofOfWork(block)

	//	返回满足条件的数据
	hash,nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	fmt.Printf("\n")
	return  block
}



//生成传世区块
func CreateGenesisBlock(data string) *Block {
	//return NewBlock(1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},data)
	return NewBlock(1,make([] byte,32,32),data)
}

//	序列化 => 将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	//	创建一个butter
	var result bytes.Buffer
	//创建一个编码器
	encoder := gob.NewEncoder(&result)
	//编码-->打包
	err := encoder.Encode(block)
	if err != nil{
		log.Panic(err)
	}
	return result.Bytes()
}

//	反序列化 => 区块字节数组转为区块对象
func DeserializeBlick(blockBytes []byte) *Block {
	var block Block
	var reader = bytes.NewReader(blockBytes)
	//	创建一个编码器
	decoder := gob.NewDecoder(reader)
	//	解包
	err := decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}
	return &block
}