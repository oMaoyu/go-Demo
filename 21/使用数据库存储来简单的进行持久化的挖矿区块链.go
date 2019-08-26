package _1

func BlockChainDemoRun() {
	block := NewBlockChain()
	cli := CLI{bc:block}
	cli.run()
}

