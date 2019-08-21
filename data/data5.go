package data

import "fmt"

type stackNode struct {
	Data interface{}
	Next *stackNode
}


// 快速创建 栈线链
func (node *stackNode) New(data ...interface{}) {
	if node == nil || data == nil {
		return
	}
	// 第一个指向空,容器指向第一个  如果有多个 一次下压  先进后出
	for _, v := range data {
		newNode := new(stackNode)
		newNode.Data = v
		newNode.Next = node.Next
		node.Next = newNode
	}
}

// 打印
func (node *stackNode) Print() {
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
func (node *stackNode) Length() int {
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
func (node *stackNode) add(index int, value interface{}, isHead bool) {
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
	newNode := new(stackNode)
	newNode.Data = value
	newNode.Next = node.Next
	node.Next = newNode
}

// 添加对应位置节点任意数据切片
func (node *stackNode) Add(index int, isHead bool, value ...interface{}) {
	for i := 0; i < len(value); i++ {
		node.add(index, value[len(value)-i-1], isHead)
	}
}

// 添加链表前置切片
func (node *stackNode) AddHead(isHead bool, value ...interface{}) {
	for i := 0; i < len(value); i++ {
		node.add(1, value[len(value)-i-1], isHead)
	}
}


// 返回对应下标节点
func (node *stackNode) popNode() *stackNode {
	if node == nil {
		return nil
	}

	return node.Next
}

// 删除单个节点
func (node *stackNode) Pop() {
	node.Next = node.popNode().ruinNode()
}

//删除指定范围开始的长度节点
func (node *stackNode) DeleteLen( len int) {
	if len-1 > node.Length() || len < 1 {
		return
	}
	for i := 0; i < len; i++{
		node.Pop()
	}
}

//删除全部节点
func (node *stackNode)DeleteAll(){
	node.DeleteLen(node.Length())
}

//销毁该节点
func (node *stackNode) ruinNode() *stackNode {
	Node := node.Next
	node.Next = nil
	node.Data = nil
	node = nil
	return Node
}
//销毁链表
func (node *stackNode) Ruin(){
	if node == nil {
		return
	}
	node.DeleteAll()
	node = nil
}




func NewStackNode() {
	node := new(stackNode)
	node.New(1,2,3,4,5,6,7,8,9,0)
	node.Print()
	//node.Pop()
	node.DeleteLen(3)
	node.Print()
}