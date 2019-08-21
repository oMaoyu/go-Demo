package data

import "fmt"

type loopNode struct {
	Data interface{}
	Next *loopNode
}

// 快速创建
func (node *loopNode) New(data ...interface{}) {
	if node == nil || data == nil {
		return
	}
	head := node
	for _, v := range data {
		newNode := new(loopNode)
		newNode.Data = v
		newNode.Next = nil
		node.Next = newNode
		node = node.Next
	}
	node.Next = head.Next
}

// 打印
func (node *loopNode) Print() {
	// 递归出口,同样避免错误
	if node == nil {
		return
	}
	start := node.Next

	for  {
		node = node.Next
		fmt.Print(node.Data, " ")
		if node.Next == start {
			break
		}
	}
	fmt.Println()
}

//  获取长度
func (node *loopNode) Length() int {
	if node == nil {
		return 0
	}
	i := 0
	start := node.Next
	for{
		i++
		node = node.Next // 修改node ，指向下一节点
		// 判断是否打印完一圈。 此时node代表尾结点
		if start == node.Next {
			break
		}

	}
	return i
}

// 插入到指定位置的节点的前后位置
func (node *loopNode) add(index int, value interface{}, isHead bool) {
	if node == nil || index <= 0 || index > node.Length() {
		return
	}
	location := 0
	if isHead {
		location = 1
	}
	start := node.Next
	for i := 0; i < index-location; i++ {
		node = node.Next
	}
	newNode := new(loopNode)
	newNode.Data = value
	newNode.Next = node.Next
	if index == node.Length() && !isHead {
		newNode.Next = start
	}
	node.Next = newNode
}

// 添加对应位置节点任意数据切片
func (node *loopNode) Add(index int, isHead bool, value ...interface{}) {
	for i := 0; i < len(value); i++ {
		node.add(index, value[len(value)-i-1], isHead)
	}
}

// 添加链表前置切片
func (node *loopNode) AddHead(isHead bool, value ...interface{}) {
	for i := 0; i < len(value); i++ {
		node.add(1, value[len(value)-i-1], isHead)
	}
}

// 添加链表尾巴切片
func (node *loopNode) AddTail(isHead bool, value ...interface{}) {
	for _, v := range value {
		node.add(node.Length(), v, isHead)
	}
}

// 返回对应下标节点
func (node *loopNode) node(index int) *loopNode {
	if node == nil || index < 0 || index > node.Length() {
		return nil
	}
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

// 删除单个节点
func (node *loopNode) Delete(index int) {
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
func (node *loopNode) DeleteLen(index, len int) {
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
func (node *loopNode)DeleteAll(){
	node.DeleteLen(1,node.Length())
}

//销毁该节点
func (node *loopNode) ruinNode() *loopNode {
	Node := node.Next
	node.Next = nil
	node.Data = nil
	node = nil
	return Node
}
//销毁链表
func (node *loopNode) Ruin(){
	if node == nil {
		return
	}
	node.DeleteAll()
	node = nil
}
func NewLoopNode(){
	node := new(loopNode)
	node.New(1,2,3,4,5,6,7,8,9,0)
	node.Print()
	fmt.Println(node.Length())
	node.add(10,213,false)
	node.Print()

	//fmt.Println()
	node.Delete(11)
	node.Print()


}
