package data

import (
	"fmt"
)

type linkNode struct {
	Data interface{}
	Next *linkNode // 只能是 本结构体指针类型。
}

// 快速创建
func (node *linkNode) New(data ...interface{}) {
	if node == nil || data == nil {
		return
	}
	for _, v := range data {
		newNode := new(linkNode)
		newNode.Data = v
		newNode.Next = nil
		node.Next = newNode
		node = node.Next
	}
}

// 打印
func (node *linkNode) Print() {
	// 递归出口,同样避免错误
	if node == nil {
		fmt.Println()
		return
	}

	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
	node.Next.Print()

}

//  获取长度
func (node *linkNode) Length() int {
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

// 插入到指定位置的节点的前后位置
func (node *linkNode) add(index int, value interface{}, isHead bool) {
	if node == nil || index <= 0 || index > node.Length() {
		return
	}
	location := 0
	if isHead {
		location = 1
	}
	for i := 0; i < index-location; i++ {
		node = node.Next
	}
	newNode := new(linkNode)
	newNode.Data = value
	newNode.Next = node.Next
	node.Next = newNode
}

// 添加对应位置节点任意数据切片
func (node *linkNode) Add(index int, isHead bool, value ...interface{}) {
	for i := 0; i < len(value); i++ {
		node.add(index, value[len(value)-i-1], isHead)
	}
}

// 添加链表前置切片
func (node *linkNode) AddHead(isHead bool, value ...interface{}) {
	for i := 0; i < len(value); i++ {
		node.add(1, value[len(value)-i-1], isHead)
	}
}

// 添加链表尾巴切片
func (node *linkNode) AddTail(isHead bool, value ...interface{}) {
	for _, v := range value {
		node.add(node.Length(), v, isHead)
	}
}

// 返回对应下标节点
func (node *linkNode) node(index int) *linkNode {
	if node == nil || index < 0 || index > node.Length() {
		return nil
	}
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

// 删除单个节点
func (node *linkNode) Delete(index int) {
	// 正常删除节点写法
	//if node == nil || index < 0 || index > node.Length() {
	//	return
	//}
	//for i := 0; i < index-1; i++ {
	//	node = node.Next
	//}
	//delectNode := node.Next
	//node.Next = delectNode.Next
	//delectNode.Next = nil
	//delectNode.Data = nil
	//delectNode = nil
	// 封装后调用写法 这里由于是对链表的学习  所以没有做正常的err机制处理
	node = node.node(index - 1)
	node.Next = node.Next.ruinNode()

}

//删除指定范围开始的长度节点
func (node *linkNode) DeleteLen(index, len int) {
	if index < 1 || index+len-1 > node.Length() || len < 1 {
		return
	}
	node = node.node(index - 1)
	Node := node.Next
	for i := 0; i < len; i++{
		Node = Node.ruinNode()
	}
	node.Next = Node
}

//删除全部节点
func (node *linkNode)DeleteAll(){
	node.DeleteLen(1,node.Length())
}

//销毁该节点
func (node *linkNode) ruinNode() *linkNode {
	Node := node.Next
	node.Next = nil
	node.Data = nil
	node = nil
	return Node
}
//销毁链表
func (node *linkNode) Ruin(){
	if node == nil {
		return
	}
	node.DeleteAll()
	node = nil
}


func NewNode() {
	node := new(linkNode)
	node.New(1, 2, 3, 4, 5, 6, 7)
	node.Print()
	node.Add(node.Length(), false, 0, 9, 8, 7)
	node.AddHead(false, "地方", "哈哈", 1, "???")
	node.AddTail(true, "432", "123", 312, "09876")
	//node.Delete(1)
	node.Print()
	//node.DeleteLen(1,4)
	//node.DeleteAll()
	node.Ruin()
	node.Print()
	fmt.Println(node)
}
