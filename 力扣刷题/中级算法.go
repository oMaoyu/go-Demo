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
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	if sum > 9 {
		return &ListNode{
			Val: sum - 10,
			Next: addTwoNumbers(addTwoNumbers(l1.Next, l2.Next), &ListNode{
				Val:  1,
				Next: nil,
			}),
		}
	} else {
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

// 3 无重复字符串子串
// 以解决
func lengthOfLongestSubstring(s string) int {
	// 思路 :  使用字典记录每个字符的位置, 如果出现重复则获取重复之间的长度  后续出现新的长度则判断长度  长度大于已有长度替换
	Len, index := 0, 0
	// 创建一个字符字典
	dict := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		// 判断是否是重复数  如果是  获取之前重复数下标 对记录下标相减  获得对应长度
		if dict[s[i]] > index {
			index = dict[s[i]]
		}
		if i-index+1 > Len {
			Len = i - index + 1
		}

		dict[s[i]] = i+1

	}
	return Len
}


//根据评论后学习所写  qweqwerq
//func lengthOfLongestSubstring(s string) int {
//
//	// num 长度   j 字符串起始下标  t 用于判断是否包含字符串
//	num, j, t := 0, 0, 0
//	for i := 0; i < len(s); i++ {
//	fmt.Println(s[j:i],string(s[i]))
//		t = strings.Index(s[j:i], string(s[i]))
//		if t == -1 {
//			if num < i-j+1 {
//				num = i - j + 1
//			}
//		}else {
//			j +=  t + 1
//		}
//		fmt.Println("i:",i)
//		fmt.Println("t:",t)
//	}
//	return num
//}

// 同样学习  跟我思路差不多  但我那个一开始比较复杂 看完后改进才变成这样 但思路是相同的
//func lengthOfLongestSubstring(s string) int {
//	val := []byte(s)
//	kvMap := make([]int, 128)
//	lens := len(s)
//	var max, num int
//	for i, j := 0, 0; i < lens && j < lens; j++ {
//		if kvMap[val[j]] > i {
//			i = kvMap[val[j]]
//		}
//		num = j - i + 1
//		if num > max {
//			max = num
//		}
//		kvMap[val[j]] = j + 1
//		fmt.Println(val[j],kvMap[val[j]],i)
//	}
//	return max
//}
//qweqwerq