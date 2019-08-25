package _0

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

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

func newBlcok(data string,prevHash []byte) *Block{
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