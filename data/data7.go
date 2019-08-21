package data

import "fmt"

//1.首先定义结构体
type MpNode struct {
	Data Dict    // 定义数据  数据为存放k v 的结构体
	Next *MpNode // 下个节点
}
type Dict struct {
	Key   string
	Value interface{}
}

//8 开始创建数组链表  也就是哈希表
// 定义数组链表
var Arr [16]*MpNode

//2 初始化
//初始化链表的头节点
func newNodeHead() *MpNode {
	node := new(MpNode)
	node.Data.Key = "头key"
	node.Data.Value = "头value"
	node.Next = nil
	return node
}

//9 初始化哈希表
func NewHash() {
	for i := 0; i < 16; i++ {
		Arr[i] = newNodeHead()
	}
}

//链方法
//3 封装结构体对象方法
// 封装key value 存放方法
func (node *MpNode) data(k string, v interface{}) *MpNode {
	if node == nil {
		node = newNodeHead()
	}
	node.Data.Key = k
	node.Data.Value = v
	return node
}

//12 根据key取value
func (node *MpNode) getKey(k string) interface{} {
	if node.Next == nil {
		return nil
	}
	for node.Next != nil {
		if node.Next.Data.Key == k {
			return node.Next.Data.Value
		} else {
			node = node.Next
		}
	}
	return nil
}

//13 根据valve取key  后面发现如果要写这个 就要给每个数组都遍历一遍后  就暂且放弃 等后续优化
//func (node *MpNode) getValue(v interface{}) string {
//	if node.Next == nil {
//		return ""
//	}
//	for node.Next != nil {
//		if node.Next.Data.Value == v {
//			return node.Next.Data.Key
//		} else {
//			node = node.Next
//		}
//	}
//	return ""
//}

//4  向该链空节点创建新数据
func (node *MpNode) add(k string, v interface{}) {
	// 首先判断k是否是该链已有的key  因为根据哈希算法简化后  同样的字符串会在同样的链里存放
	if node.getKey(k) != nil {
		return
	}
	//遍历到尾节点
	for node.Next != nil {
		node = node.Next
	}
	// 创建新的尾节点
	node.Next = node.Next.data(k, v)
}

//6 发现这样打印看不太方便 于是写个遍历值
func (node *MpNode) Log() {
	if node.Next == nil {
		return
	}
	fmt.Println(node.Next.Data)
	node.Next.Log()
}

// 7计算链表长度 防止数据堆积 为后续优化做准备
func (node *MpNode) Length() int {
	if node == nil {
		return 0
	}
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
	}
	return i
}

//哈希表方法
//10 往表里存key值
func SetKey(k string, v interface{}) {
	// 存k先判断存哪里  这里才使用了散列算法计算
	num := hashNum(k)
	Arr[num].add(k, v)
}
//14 取key值
func GetKey(k string) interface{} {
	num := hashNum(k)
	return Arr[num].getKey(k)
}

//11 获取16以内的散列算法
func hashNum(key string) int {
	var index int = 0
	index = int(key[0])
	// 询价累积新的数值
	for k := 0; k < len(key); k++ {
		index *= (1103515245 + int(key[k]))
	}
	// 右位移2^27次方  这里可以位移任何数  只要别太大导致直接清空
	index >>= 27
	// 按位&  这里想要求32以内就 32-1即可   也就是说  index跟 1111 进行&  得到15以内的数  跟11111取余则得31以内的数
	index &= 16 - 1

	return index
}

//5 测试下添加数据是否有问题
func MpDemo() {
	//node := newNodeHead()
	//node.add("1", "2")
	//node.add("2", 3)
	//node.Log()
	//fmt.Println(node.Length())
}
