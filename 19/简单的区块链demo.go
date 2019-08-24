package _9


func BlockChainDemoRun() {
	block := NewBlockChain()
	block.AddBlock("这是第一条区块链数据")
	block.AddBlock("这是第二条区块链数据")
	block.Log()

}
