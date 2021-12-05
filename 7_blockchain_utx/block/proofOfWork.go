package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)
//	难度	256位Hash里面前面至少要有16个零
const targetBigs = 8
type  ProofOfWork struct {
	Block *Block	//	当前要验证的区块
	Target *big.Int	//	大数据存储 目标hash
}

//	校样hash是否有效
func (pow *ProofOfWork)IsValid() bool {
	hashInt := new(big.Int)
	hashInt.SetBytes(pow.Block.Hash)

	return pow.Target.Cmp(hashInt) == 1
}
func (p *ProofOfWork) Run() ([]byte, int64) {
	//	把Block数据(除了hash)平接成字节数组

	//	生成hash

	//判断hash有效性，如果满足条件，跳出循环
	 nonce  := 0
	 hashInt := new(big.Int)	//	存储我们生成的hash
	var hash [32]byte
	for  {
		//	准备数据
		dataBytes := p.prepareData(nonce)

		//	生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("生成的哈希：\r %x ",hash)
		//	将hash存储到hashInt中
		hashInt.SetBytes(hash[:])
		//	判断hashInt是否小于Block里面的target
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//	p.target > hashInt
		if p.Target.Cmp( hashInt) == 1 {
			break
		}
		nonce += 1
	}
	return  hash[:], int64(nonce)
}

//数据平接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte  {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.HashTransactions(),
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBigs)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)
	return  data
}
//	创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	//	创建一个初始值为1的target
	target := big.NewInt(1)
	//	左移 256 - targetBit
	target = target.Lsh(target, uint(256 - targetBigs))

	pow := &ProofOfWork{block, target}
	return  pow
}