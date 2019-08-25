package _0

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	//提供数据的区块
	block *Block
	// 对应难度值
	target *big.Int
}

// new一个难度值
func NewProof(block *Block) *ProofOfWork {
	// 添加对应block数据
	pow := ProofOfWork{block: block}
	// 添加难度值
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	bigTmp := big.Int{}
	bigTmp.SetString(targetStr, 16)
	pow.target = &bigTmp
	return &pow
}

// 拼接当前数据字段,根据随机数值来组合新数据进行哈希
func (pow *ProofOfWork) PrepareData(nonce uint64) []byte {
	b := pow.block

	byteTmp := [][]byte{
		uintToByte(b.Version),
		b.PrevHash,
		b.MerKleRoot,
		uintToByte(b.TimeStamp),
		uintToByte(b.Bits),
		uintToByte(nonce),
		b.Data,
	}
	return bytes.Join(byteTmp, []byte{})
}

// 开始运行进行挖矿
//返回1 为挖矿挖到的区块哈希
//返回2 为对应哈希的随机数
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	// 定义随机数变量 为初始值0
	var nonce uint64

	fmt.Println("开始挖矿....")
	var currHash [32]byte

	for {
		fmt.Printf("hash:%x\r", currHash)
		//1 给数据拼接在一起
		data := pow.PrepareData(nonce)
		//2 计算出哈希值
		currHash = sha256.Sum256(data)
		//3 给哈希值进行转换来进行比较  转换类型为big.int
		tmp := big.Int{}
		tmp.SetBytes(currHash[:])

		//跟自己当前本身区块难度值进行比较  如果小则挖矿成功  小为-1
		if tmp.Cmp(pow.target) == -1 {
			//当前哈希值小于难度值
			fmt.Printf("挖矿成功, 当前的哈希值:%x, nonce : %d\n", currHash[:], nonce)
			break
		} else {
			// 哈希不小于  随机数自增
			nonce++
		}
	}
	return currHash[:], nonce
}
// 当挖矿结束后  其他节点进行对这次挖矿的运算是否满足
//验证这次是否有效
func (pow *ProofOfWork)IsValidBlock() bool {
	// 获取对应区块对应的随机值 并拼接数据
	data := pow.PrepareData(pow.block.Nonce)
	// 对数据进行哈希计算
	currHash := sha256.Sum256(data)
	// 给哈希进行转换
	tmp := big.Int{}
	tmp.SetBytes(currHash[:])

	return tmp.Cmp(pow.target) == -1
}