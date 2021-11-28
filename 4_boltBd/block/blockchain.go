package block

//	区块链结构体
type Blockchain struct {
	Blocks []*Block //	存储有序区块
}
//	创建带有传世区块的区块链
func CreateBlockchainGenesisBlock(data string) *Blockchain {
	//	创建传世区块
	genesisBlock := CreateGenesisBlock(data)
	//	返回区块对象
	return &Blockchain{[]*Block{genesisBlock}}
}

//	添加区块到区块链中
func (bc *Blockchain) AddBlockToBlockchain(height int64, preHash []byte, data string)  {
	//	创建一个新的区块
	newBlock := NewBlock(height, preHash, data)
	//	往数区块链数组中添加新区块
	bc.Blocks = append(bc.Blocks, newBlock)
}