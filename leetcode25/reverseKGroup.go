package leetcode25

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/* 看了两小时才再自己写...第二天在看一次，啊不会啊.*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	fmt.Println("reverseKGroup begin")

	var result ListNode
	var cur = &result
	cur.Next = head

	var temp = head
	var next *ListNode
	i := 1
	for temp != nil && temp.Next != nil {
		if i%k != 0 {
			next = temp.Next
			temp.Next = next.Next
			next.Next = cur.Next
			cur.Next = next
			i++
		} else {
			cur = temp
			temp = temp.Next
			i = 1
		}
	}
	if i%k != 0 {
		temp = cur.Next
		for temp != nil && temp.Next != nil {
			next = temp.Next
			temp.Next = next.Next
			next.Next = cur.Next
			cur.Next = next
		}
	}
	return result.Next
}
