package _1

import (
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}

// 命令使用说明
const Usage = `
	./oMaoyu add <需要添加的数据>  ==>添加区块
	./oMaoyu log    ==>打印区块链
`

func (cli *CLI) run() {
	arr := os.Args
	if len(arr) < 2    {
		fmt.Println("输入参数无效，请检查！")
		fmt.Println(Usage)
		return
	}
	fmt.Println(arr)
	switch arr[1] {
	case "add":
		fmt.Println("添加区块命令被调用!")
		if len(arr) <3 {
			fmt.Println("输入参数无效，请检查！")
			fmt.Println(Usage)
			return
		}
		cli.Add(arr[2])
	case "log":
		cli.Log()
	default:
		fmt.Println("输入参数无效，请检查！")
		fmt.Println(Usage)
	}
}
func (cli *CLI) Log() {
	cli.bc.Log()
}

func (cli *CLI) Add(data string) {
	cli.bc.add(data)
}
