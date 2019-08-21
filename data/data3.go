package data

import "fmt"

type TwoNode struct {
	Data interface{}
	Next *TwoNode // 只能是 本结构体指针类型。
	Last *TwoNode
}
func (node *TwoNode)New(data ...interface{}){
	if node == nil || data == nil {
		return
	}
	for _,v := range data {
		newNode := new(TwoNode)
		newNode.Last = node
		newNode.Data = v
		newNode.Next = nil
		node.Next = newNode
		node = newNode
	}
}

// 打印
func (node *TwoNode) Print() {
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
func (node *TwoNode) Print1()  {
	if node == nil {
		return
	}
	for node.Next != nil{
		node  = node.Next
	}
	for node.Last  != nil {
		fmt.Print(node.Data, " ")
		node = node.Last
	}

}

//  获取长度
func (node *TwoNode) Length() int {
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
func (node *TwoNode) add(index int, value interface{}, isHead bool) {
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
	newNode := new(TwoNode)
	newNode.Data = value
	newNode.Last = node
	newNode.Next = node.Next
	if node.Next != nil {
		node.Next.Last = newNode
	}
	node.Next = newNode
}

// 添加对应位置节点任意数据切片
func (node *TwoNode) Add(index int, isHead bool, value ...interface{}) {
	for i := 0; i < len(value); i++ {
		node.add(index, value[len(value)-i-1], isHead)
	}
}

// 添加链表前置切片
func (node *TwoNode) AddHead(isHead bool, value ...interface{}) {
	for i := 0; i < len(value); i++ {
		node.add(1, value[len(value)-i-1], isHead)
	}
}

// 添加链表尾巴切片
func (node *TwoNode) AddTail(isHead bool, value ...interface{}) {
	for _, v := range value {
		node.add(node.Length(), v, isHead)
	}
}

// 返回对应下标节点
func (node *TwoNode) node(index int) *TwoNode {
	if node == nil || index < 0 || index > node.Length() {
		return nil
	}
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

// 删除单个节点
func (node *TwoNode) Delete(index int) {

	node = node.node(index - 1)
	node.Next = node.Next.ruinNode()
	node.Next.Last = node
}

//删除指定范围开始的长度节点
func (node *TwoNode) DeleteLen(index, len int) {
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
func (node *TwoNode)DeleteAll(){
	node.DeleteLen(1,node.Length())
}

//销毁该节点
func (node *TwoNode) ruinNode() *TwoNode {
	Node := node.Next
	node.Last = nil
	node.Next = nil
	node.Data = nil
	node = nil
	return Node
}
//销毁链表
func (node *TwoNode) Ruin(){
	if node == nil {
		return
	}
	node.DeleteAll()
	node = nil
}


func NewTwoNode(){
	node := new(TwoNode)
	node.New(1,2,3,4,5,6,7,8,9,0)
	node.Print()
	//node.Print1()
	fmt.Println(node.Length())
	node.add(10,213,false)
	node.Print1()
	fmt.Println()
	node.Delete(3)
	node.Print1()
}
