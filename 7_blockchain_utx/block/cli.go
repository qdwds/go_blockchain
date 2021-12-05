package block

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct {
	//Blockchain *Blockchain
}

func printUsage()  {
	fmt.Println("genesis -data 创建创世区块")
	fmt.Println("addblock -data DATA -- 交易数据")
	fmt.Println("printchain -- 输出区块信息")
}
func isValidArgs()  {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
//	添加区块
func (cli *Cli)addBlock(txs []*Transaction)  {
	if DBExists() == false{
		fmt.Println("请先创建创世区块！")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.AddBlockToBlockchain(txs)
}
//	输出
func (cli *Cli) printChain() {
	if DBExists() == false{
		fmt.Println("请先创建创世区块！")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.Printchain()
}

//	创建创世区块
func (cli *Cli) createGenesis(txs []*Transaction)  {
	CreateBlockchainGenesisBlock(txs)
}
func (cli *Cli) Run()  {
	//	通过flag 来调用区块链；
	addBlockCmd := flag.NewFlagSet("addlock",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)
	genesisCmd := flag.NewFlagSet("genesis",flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data","https://","添加新区块")
	flagCreateGenesis := genesisCmd.String("data","genesis data ..","创建创世区块")

	isValidArgs()
	fmt.Println(os.Args[1],"os.Argsos.Argsos.Args")
	switch os.Args[1] {
		case "addblock":
			err := addBlockCmd.Parse(os.Args[2:])
			if err != nil{
				log.Panic(err)
			}
		case "printchain":
			err := printChainCmd.Parse(os.Args[2:])
			if err != nil{
				log.Panic(err)
			}
		case "genesis":
			err := genesisCmd.Parse(os.Args[2:])
			if err != nil{
				log.Panic(err)
			}
		default:
			printUsage()
	}

	if addBlockCmd.Parsed(){
		if *flagAddBlockData == ""{
			printUsage()
			os.Exit(1)
		}
		//fmt.Println(*flagAddBlockData)
		cli.addBlock(*flagAddBlockData)
	}
	if printChainCmd.Parsed(){
		fmt.Println("输出所有区块数据")
		cli.printChain()
	}

	if genesisCmd.Parsed() {
		if *flagCreateGenesis == "" {
			fmt.Println("交易数据不能为空")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesis(*flagCreateGenesis)
	}
}

