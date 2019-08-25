package _0

func BlockChainDemoRun() {
	block := NewBlockChain()
	block.add("这是第一条区块链数据")
	block.add("这是第二条区块链数据")
	block.Log()

}

