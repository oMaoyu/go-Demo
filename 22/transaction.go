package _1

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

// 首先比特币交易需要核实每一笔的交易信息
// 需要追究这笔钱的头   所以必要参数为
// 交易输出
// 交易输入
// 时间戳
// 交易的id
type Transaction struct {
	Id   []byte //id
	Time int64  // 时间戳
	// 而比特币的交易里  本次的输入就是前一个的输出  每一个交易的块里,不代表只有一个用户在进行交易  所以输入输出应该为切片
	Inputs []Input
	Outputs []Output
}

// 交易输入
//每一笔交易都拥有一个输入脚本  也就是解锁脚本
//解锁脚本描述了某人可以花费这笔前，提供的是私钥签名
type Input struct {
	I_id       []byte // 所引用的输出id
	I_vout     int32 // 所引用的输出的索引
	I_scripsig string // 解锁脚本,描述付款人
}

// 交易输出
//每一笔输出都会有一个输出脚本
//输出脚本又叫做“锁定脚本”，描述一笔钱属于哪个地址(用收款的地址的公钥哈希锁定)
type Output struct {
	Value        float32 // 转账金额
	O_scriptlock string  // 锁定脚本,描述负付款人
}

// 生成交易ID  对交易体进行编码后哈希
func (tx *Transaction) setId() {
	var buff bytes.Buffer
	en := gob.NewEncoder(&buff)
	err := en.Encode(tx)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(buff.Bytes())
	tx.Id = hash[:]
}
// 设置当前挖矿奖励
const reward = 50.0

// 比特币中存在两种形式的交易
// 1 挖矿交易  特点 没有输入 只有输出

// 挖矿成功后开始交易奖励
// 在比特币当中  所有矿工有权写一句话到coinbase当中
// 现状通常是矿池写自己矿池的名字

//data: 写入挖矿交易的信息
//miner: 是挖矿人
func NewCoinbaseTx(data string,miner string) *Transaction{

	// 创建输入
	input := Input{
		I_id:nil,
		I_vout: -1,
		I_scripsig:data,
	}
	// 创建输出
	output := Output{
		Value: reward,
		O_scriptlock:miner,
	}
	tx := Transaction{
		Id:nil,
		Inputs:[]Input{input},
		Outputs:[]Output{output},
		Time:time.Now().Unix(),
	}

	tx.setId()

	return &tx
}



// 2 普通转账交易