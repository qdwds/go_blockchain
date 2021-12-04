package block

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct {
	Blockchain *Blockchain
}

func printUsage()  {
	fmt.Println("Usage:")
	fmt.Println("\taddblock - data DATA -- 交易数据")
	fmt.Println("\tprintchain -- 输出区块信息")
}
func isValidArgs()  {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
//	添加区块
func (cli *Cli)addBlock(data string)  {
	cli.Blockchain.AddBlockToBlockchain(data)
}
//
func (cli *Cli) printChain() {
	cli.Blockchain.Printchain()
}
func (cli *Cli) Run()  {
	//	通过flag 来调用区块链；
	addBlockCmd := flag.NewFlagSet("addlock",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)
	flagAddBlockData := addBlockCmd.String("data","https://","data___")
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
}