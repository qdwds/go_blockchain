package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//	区块结构体
type Block struct {
	Height int64			//	区块高度
	PrevBlockHash []byte	//	上一个区块的HASH
	Data []byte				//	交易数据
	Timestamp int64			//	时间戳
	Hash []byte				//	当前HASH
}

//	创建新 区块
func NewBlock( height int64, PrevBlockHash []byte,data string) *Block {
  	block := &Block{height, PrevBlockHash, []byte(data),time.Now().Unix(), nil}
  	block.SetHash()
	return  block
}

//	区块信息 => hash
func (b *Block)SetHash()  {
	//	height
	heightBytes := IntToHex(b.Height)
	//	time
	timeString := []byte(strconv.FormatInt(b.Timestamp,10))
	//	将所有信息拼接
	headers := bytes.Join([][]byte{heightBytes,b.PrevBlockHash,b.Data,timeString,b.Hash},[]byte{})
	//	转成hash
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//生成传世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},data)
}