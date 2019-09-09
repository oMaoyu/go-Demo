package likou

import (
	"fmt"
)

func (node *ListNode) Log() {
	if node != nil {
		fmt.Println(node.Val)
		node.Next.Log()
	}
}

//2 两数相加 初版  正序  过了

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	if sum > 9 {
		return &ListNode{
			Val:  sum - 10,
			Next:addTwoNumbers( addTwoNumbers(l1.Next, l2.Next),&ListNode{
				Val:  1,
				Next: nil,
			}),
		}
	}else {
		return &ListNode{
			Val:  sum,
			Next: addTwoNumbers(l1.Next, l2.Next),
		}
	}

}

// 逆序  通过

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil{
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//
//	sum := l1.Val + l2.Val
//	node := addTwoNumbers(l1.Next,l2.Next)
//	if sum > 9{
//
//		// 创建一个新的node
//		newNode := &ListNode{
//			Val:  1,
//			Next: nil,
//		}
//		return &ListNode{
//			Val:  sum - 10,
//			Next: addTwoNumbers(node, newNode),
//		}
//	}else{
//		return &ListNode{
//			Val:  sum,
//			Next: node,
//		}
//	}
//
//}