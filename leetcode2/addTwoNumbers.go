package leetcode2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head = ListNode{}
	var curr = &head

	var list1 = l1
	var list2 = l2

	var temp = 0
	for list1 != nil || list2 != nil {
		var x = 0
		if list1 != nil {
			x = list1.Val
			list1 = list1.Next
		}

		var y = 0
		if list2 != nil {
			y = list2.Val
			list2 = list2.Next
		}

		var sum = temp + x + y
		temp = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
	}
	if temp > 0 {
		curr.Next = &ListNode{temp, nil}
		curr = curr.Next
	}
	return head.Next
}
