package leetcode21

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*这是一个第一次提交直接4ms 超过100%的人的代码，感动！！！连自测都不需要*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println("mergeTwoLists begin")
	dummyHead := &ListNode{}
	temp := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			temp.Next = l2
			temp = temp.Next
			l2 = l2.Next
		} else {
			temp.Next = l1
			temp = temp.Next
			l1 = l1.Next
		}
	}
	if l1 == nil {
		temp.Next = l2
	}
	if l2 == nil {
		temp.Next = l1
	}
	return dummyHead.Next
}
