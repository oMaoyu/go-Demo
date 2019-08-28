package _1

import (
	"fmt"
	"github.com/boltdb/bolt"
)

// 定义区块链
// 此时区块链定义对象应该为db服务器对象
//最后区块的哈希值  应为是以最后区块哈希值作为key去写
type BlockChain struct {
	//blocks []*Block
	db   *bolt.DB
	tail []byte
}

const dataName = "oMaoyu blockChain demo name :oMaoyu"
const blockchainname = "./22/blockchain.db"
const blockcbucket = "blockbucket"
const lastBLockHashKey = "lastBLockHashKey"

func NewBlockChain() *BlockChain {

	//创建区块链时，向里面写入一个创世区块
	//包含两个功能：
	//1. 如果区块链不存在，则创建，写入创世块
	db, err := bolt.Open(blockchainname, 0600, nil)
	if err != nil {
		panic(err)
	}
	var lastHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		//获取bucket
		b := tx.Bucket([]byte(blockcbucket))
		if b == nil {
			//创建bucket
			b, err = tx.CreateBucket([]byte(blockcbucket))
			if err != nil {
				return err
			}
			// 创建挖矿交易coinbase
			coin := NewCoinbaseTx(dataName,"oMaoyu")

			//写入创世块
			block := newBlcok([]*Transaction{coin}, nil)
			//写入数据库
			//给当前区块的哈希当做唯一Key去存储  给区块进行转换
			b.Put(block.Hash, block.ToByte())
			b.Put([]byte(lastBLockHashKey), block.Hash)
			//设置lastBlockHashKey的值
			lastHash = block.Hash
		} else {
			//2. 如果区块链存在， 直接读取最后一个区块的哈希
			lastHash = b.Get([]byte(lastBLockHashKey))
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	return &BlockChain{db: db, tail: lastHash}
}

func (b *BlockChain) add(txs []*Transaction) {

	newBlock := newBlcok(txs, b.tail)
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockcbucket))
		if bucket == nil {
			panic("数据库为空,无法添加")
		}
		bucket.Put(newBlock.Hash, newBlock.ToByte())
		bucket.Put([]byte(lastBLockHashKey), newBlock.Hash)
		b.tail = newBlock.Hash
		return nil
	})
	if err != nil {
		panic(err)
	}
}
func (bc *BlockChain) Log() {
	prev := bc.tail
	for prev != nil {
		block :=bc.next(prev)
		block.print()
		prev = block.PrevHash
	}
}

type OutputInfo struct {
	Out Output
	Id []byte
	Index int32
}

// 遍历账本 找到某个地址的所有utxo
func  (bc *BlockChain) FindMyUtxo(add string)[]OutputInfo{
	var out []OutputInfo
	//定义一个map，用于存储已经消耗过的output，在遍历inputs是动态更新
	//key ==> 交易id
	//value ==> 1
	tmpOut := make(map[string][]int32)
	// 开始遍历block
	prev := bc.tail
	for prev != nil {
		block :=bc.next(prev)
		prev = block.PrevHash
		// 遍历交易
		for _,tx :=range block.Transaction {
			to:
			//遍历输出 判断当前输出的地址是否是自己查找的地址
			for index,output := range tx.Outputs{
				if output.O_scriptlock == add {
					// 过滤
					// 当前交易id
					id := string(tx.Id)
					// 判断当前交易id是否被消耗过
					indexArr := tmpOut[id]
					if len(indexArr) != 0  {
						// 说明当前的交易是被消耗过的
						for _,ind := range indexArr{
							if ind == int32(index) {
								continue to
							}
						}
					}
					fmt.Printf("找到了属于'%s'的output: %d\n", add, index)
					fmt.Println(output.Value)
					out = append(out, OutputInfo{output,tx.Id,int32(index)})
				}
			}
			// 遍历输入  表示出和自己相关的输出  就是已经使用过的
			for _,input := range tx.Inputs{
				if input.I_scripsig == add {
					tmpOut[string(input.I_id)] = append(tmpOut[string(input.I_id)],input.I_vout)
				}
			}
			
		}
	}
	return out
}

// 根据账本去寻找满足条件的输出,拼接成字典返回
func (bc *BlockChain) findNeedUtxoinfos(from string, amount float32) (map[string][]int32, float32) {
	// 初始化返回值
	tmpMap := make(map[string][]int32)
	var tmpValue float32
	//首先遍历该地址的所有历史交易  找到能用的数据
	utxoInfos := bc.FindMyUtxo(from)

	// 选择满足条件的数据进行返回
	for _,info := range utxoInfos{
		tmpValue += info.Out.Value
		key := string(info.Id)
		tmpMap[key] = append(tmpMap[key],info.Index)
		// 如果现有金额满足消费金额  不继续遍历 直接返回
		if tmpValue > amount{
			 break
		}
	}
	return tmpMap,tmpValue
}


func (bc *BlockChain) next(hash []byte)(block *Block) {
	err := bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockcbucket))
		if bucket == nil{
			panic("数据库为空,无法遍历")
		}
		data := bucket.Get(hash)
		block = ToBlock(data)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return block
}
