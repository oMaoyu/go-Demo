package _9

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

/*
区块链结构包含字段
1.当前区块哈希
2.前区块哈希
3.数据
4.版本号
5.梅克尔根
6.时间戳
7.难度值
8.随机数
*/
//定义区块结构
type Block struct {
	//1. 版本号
	Version uint64

	//2. 前哈希
	PrevHash []byte

	//3. 当前区块哈希
	Hash []byte

	//4. 梅克尔根
	MerKleRoot []byte

	//5. 时间戳
	TimeStamp uint64 //使用 date +%s 来查看

	//6. 难度值：Bits
	Bits uint64 //可以推算出当前网络的难度值

	//7. 随机数
	Nonce uint64

	//8. 交易数据
	Data []byte
}

//提供创建区块的方法  对创建方法进行简单封装
func NewBlock(data string, prevHash []byte) *Block {
	b := &Block{
		Version:    0,
		PrevHash:   prevHash,
		Hash:       nil,
		MerKleRoot: nil,
		TimeStamp:  uint64(time.Now().Unix()),
		Bits:       0,
		Nonce:      0,
		Data:       []byte(data),
	}
	b.SetHash()
	return b
}

// 打印区块内容
func (b *Block) print() {
	fmt.Printf("Version: %d\n", b.Version)
	fmt.Printf("PrevHash: %x\n", b.PrevHash)
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Printf("MerKleRoot: %x\n", b.MerKleRoot)
	fmt.Printf("TimeStamp: %d\n", b.TimeStamp)
	fmt.Printf("Bits: %d\n", b.Bits)
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Println()
}

// 计算该区块参数的hash封装
func (b *Block) SetHash() {
	byteTmp := [][]byte{
		b.PrevHash,
		b.Hash,
		b.Data,
		uintToByte(b.Version),
		b.MerKleRoot,
		uintToByte(b.TimeStamp),
		uintToByte(b.Bits),
		uintToByte(b.Nonce),
	}
	data := bytes.Join(byteTmp, []byte{})

	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}
