package likou

import (
	"fmt"
	"strings"
)

func Run() {
	//fmt.Println(twoSum([]int{1, 4, 5, 8,9,13,14}, 18))
	//fmt.Println(numJewelsInStones("aD","gacDDADkew"))
	//fmt.Println(defangIPaddr("172.0.0.1"))
	//fmt.Println(removeOuterParentheses("(()())(())"))

	node := &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	node2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	addTwoNumbers(node,node2).Log()
}

func twoSum(numbers []int, target int) []int {
	s := 0
	e := len(numbers) - 1
	for s < e {
		sum := numbers[s] + numbers[e]
		if sum == target {
			break
		} else if sum > target {
			e--
		} else {
			s++
		}
	}
	return []int{s + 1, e + 1}
}

// 771 宝石
func numJewelsInStones(J string, S string) int {
	str := S
	for _, j := range []byte(J) {
		str = strings.Replace(str, string(j), "", -1)
	}
	fmt.Println(str)
	return len(S) - len(str)
}

// 1108 ip无效化
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

// 237 删除链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	n := node.Next
	node.Val = n.Val
	node.Next = n.Next
	n = nil
}
// 938  二叉树搜索和
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root  == nil {
		return 0
	}
	if root.Val >= L && root.Val <= R || root.Val >= R && root.Val <= L {
		return rangeSumBST(root.Left,L,R) + rangeSumBST(root.Right,L,R) + root.Val
	}else {
		return rangeSumBST(root.Left,L,R) + rangeSumBST(root.Right,L,R)
	}

}
// 1021 删除括号
func removeOuterParentheses(S string) string {
	data := ""
	l := 0
	index := 1
	for k,s := range S{
		if string(s) == "(" {
			l ++
		}
		if string(s) == ")"{
			l --
		}
		if l == 0 {
			data += S[index:k]
			index = k+2
		}
	}
	return data
}
// 709 大小写转换
func toLowerCase(str string) string {
	lowerStr := ""
	for _,s := range str{
		if s >= 65 && s <= (65+24) {
			lowerStr += string(s+32)
		}else {
			lowerStr += string(s)
		}
	}

	return lowerStr
}
//617 合并二叉树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// 如果t1 为空返回t2
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left,t2.Left)
	t1.Right = mergeTrees(t1.Right,t2.Right)


	return t1
}