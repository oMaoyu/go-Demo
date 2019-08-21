package data

import (
	"fmt"
	"reflect"
)

type Binary struct {
	Data interface{}
	left *Binary
	right *Binary
}

// 创建二叉树 ——  目前打印下来 全是左节点
var i = -1
func (node *Binary) Create(data []interface{})*Binary {
	if node == nil {
		return nil
	}
	i = i +1
	// 创建二叉树结点
	var bin *Binary
	if i >= len(data) {
		return nil
	}else {
		bin = new(Binary)
		bin.Data = data[i]
		bin.left = bin.Create(data)
		bin.right = bin.Create(data)
	}
	return bin
}

func (node *Binary)New(index int,data []interface{})*Binary{
	if node == nil {
		return nil
	}
	bin := &Binary{data[index],nil,nil}
	// 设置完全二叉树左节点  其特征是深度 *2+1为左节点  +2为右节点
	if index< len(data) && 2*index+1 < len(data){
		fmt.Println()
		bin.left = bin.New(index*2+1,data)
	}
	if i< len(data) && 2*index+2 < len(data){
		bin.right = bin.New(index*2+2,data)
	}
	return bin

}

// 遍历二叉树 —— 先序遍历： DLR
func (node *Binary) PrevSort() {
	// 递归出口
	if node == nil {
		return
	}
	// 先D, 先打印数据域
	fmt.Print(node.Data, " ")
	// 再左, 左子树递归调用本函数
	node.left.PrevSort()
	// 再右，右子树递归调用本函数
	node.right.PrevSort()
}

// 遍历二叉树 —— 中序遍历： LDR
func (node *Binary) MidSort() {
	if node == nil {
		return
	}
	// 先左，左子树递归调用本函数
	node.left.MidSort()
	// 再中, 打印数据域
	fmt.Print(node.Data, " ")
	// 再右, 右子树递归调用本函数
	node.right.MidSort()
}

// 遍历二叉树 —— 后序遍历： LRD   --- 73415620
func (node *Binary) PostSort() {
	if node == nil {
		return
	}
	// 先左， 左递归
	node.left.PostSort()
	// 再右， 右递归
	node.right.PostSort()
	// 再中, 打印
	fmt.Print(node.Data, " ")
}

// 获取二叉树的高度（深度）
func (node *Binary) Height() int {
	// 递归出口
	if node == nil {
		return 0
	}
	// 左子树递归进入
	lh := node.left.Height()
	// 右子树递归进入
	rh := node.right.Height()
	// 累加递归的返回值，左子树  和 右子树 比较
	if lh > rh {
		lh++
		return lh
	} else {
		rh++
		return rh
	}
}

// 统计二叉树的叶子结点个数  --- 全局变量
var num = 0 // 记录叶子数
func (node *Binary) LeafNum() {
	// 递归出口
	if node == nil {
		return
	}
	// 找寻叶子结点。 左子树 == 右子树 == nil
	if node.left == nil && node.right == nil {
		num++
	}
	// 左右子树各自递归调用本函数
	node.left.LeafNum()
	node.right.LeafNum()
}

// 统计二叉树的叶子结点个数  --- 指针传参
func (node *Binary) LeafNum2(n *int) {
	// 递归出口
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		(*n)++
	}
	// 左右子树各自递归调用本函数
	node.left.LeafNum2(n)
	node.right.LeafNum2(n)
}

// 查找是否存在数据
func (node *Binary) Search(Data interface{}) {
	if node == nil {
		return
	}
	// 比较类型、数值
	if reflect.DeepEqual(node.Data, Data) {
		fmt.Println("存在！")
		return
	}
	// 左右子树各自递归调用本函数
	node.left.Search(Data)
	node.right.Search(Data)
}

// 销毁二叉树
func (node *Binary) Destroy() {
	if node == nil {
		return
	}
	// 左子树递归调用本函数， 销毁左子树
	node.left.Destroy()
	node.left = nil
	// 右子树递归调用本函数， 右毁左子树
	node.right.Destroy()
	node.right = nil
	node = nil
}

// 二叉树的翻转
func (node *Binary) Reverse() {
	if node == nil {
		return
	}
	// 利用Go特有多重赋值，交换左、右子树
	node.left, node.right = node.right, node.left

	// 递归
	node.left.Reverse()
	node.right.Reverse()
}

// 二叉树的拷贝
func (node *Binary) Copy() *Binary {
	if node == nil {
		return nil
	}
	// 左子树递归进入。
	lChild := node.left.Copy()

	// 右子树递归进入。
	rChild := node.right.Copy()

	// 创建新结点
	newNode := new(Binary)
	newNode.Data = node.Data
	newNode.left = lChild
	newNode.right = rChild

	return newNode
}
func NewBinary(){
	node := new(Binary)
	node = node.New(0,[]interface{}{1,2,3,4,5,6,7,8,9})
	node.PrevSort()
}