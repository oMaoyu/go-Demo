package _0

import "fmt"

// 定义区块链
type BlockChain struct {
	blocks []*Block
}

const dataName = "oMaoyu blockChain demo name :oMaoyu"

func NewBlockChain() *BlockChain {
	block := newBlcok(dataName, nil)
	return &BlockChain{blocks: []*Block{block}}
}

func (b *BlockChain) add(data string) {
	block := b.blocks[len(b.blocks)-1]
	newBlock := newBlcok(data, block.Hash)
	b.blocks = append(b.blocks, newBlock)
}
func (bc *BlockChain) Log() {
	for _, data := range bc.blocks {
		proof := NewProof(data)
		proof.block.Hash, proof.block.Nonce = proof.Run()
		fmt.Println(proof.IsValidBlock())
		data.print()

	}
}
