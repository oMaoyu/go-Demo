package _9

// 定义区块链
type BlockChain struct {
	blocks []*Block
}

// 设置第一个区块的值  也就是创世区块的值
const genesisInfo = "oMaoyu demo 的创世区块值 "

// 创建区块链
func NewBlockChain() *BlockChain {
	block := NewBlock(genesisInfo, nil)
	return &BlockChain{blocks: []*Block{block}}
}

// 添加新的区块链
func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
func (bc *BlockChain)Log(){
	for _,data := range bc.blocks{
		data.print()
	}
}
