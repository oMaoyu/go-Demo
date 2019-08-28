package _1

import (
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	bc *BlockChain
}

// 命令使用说明
const Usage = `
	./oMaoyu add <需要添加的数据>  ==>添加区块
	./oMaoyu log    ==>打印区块链
	./oMaoyu utxo  <需要添加的数据>  ==> 矿工名
	./oMaoyu send <FROM> <TO> <AMOUNT> <MINER> <DATA>  转账人  被转账人  转账金额  挖矿人 挖矿信息
`

func (cli *CLI) run() {
	arr := os.Args
	if len(arr) < 2 {
		fmt.Println("输入参数无效，请检查！")
		fmt.Println(Usage)
		return
	}
	fmt.Println(arr)
	switch arr[1] {
	case "add":
		fmt.Println("添加区块命令被调用!")
		if len(arr) < 3 {
			fmt.Println("输入参数无效，请检查！")
			fmt.Println(Usage)
			return
		}
		cli.Add(arr[2])
	case "log":
		cli.Log()
	case "utxo":
		if len(arr) < 3 {
			fmt.Println("输入参数无效，请检查！")
			fmt.Println(Usage)
			return
		}
		cli.Utxo(arr[2])
	case "send":
		//./blockchain send <FROM> <TO> <AMOUNT> <MINER> <DATA> fmt.Println("send命令被调用!")
		if len(arr) != 7 {
			fmt.Println("输入参数无效，请检查!")
			fmt.Println(Usage)
			return
		}
		from := arr[2]
		to := arr[3]
		amountStr := arr[4]
		amount, _ := strconv.ParseFloat(amountStr, 32)
		miner := arr[5]
		data := arr[6]
		cli.Send(from, to, float32(amount), miner, data)
	default:
		fmt.Println("输入参数无效，请检查！")
		fmt.Println(Usage)
	}
}
func (cli *CLI) Log() {
	cli.bc.Log()
}

func (cli *CLI) Add(data string) {
	//cli.bc.add()
}
func (cli *CLI) Utxo(data string) {
	cli.bc.FindMyUtxo(data)
}
// 比特币当中任何的交易都是输出组  只有在矿工挖矿成功的那一瞬间才会被记录下来
func (cli *CLI) Send(from, to string, amount float32, miner string, data string) () {
	fmt.Printf("%s 向 %s 转账 %f，由 %s 记录!", from, to, amount, miner)
	tx := NewCoinbaseTx(data,miner)
	txs := []*Transaction{tx}
	tx,err := NewTransaction(from,to,amount,cli.bc)
	if err != nil {
		fmt.Println(err)
	}else {
		txs = append(txs, tx)
	}
	cli.bc.add(txs)
}
