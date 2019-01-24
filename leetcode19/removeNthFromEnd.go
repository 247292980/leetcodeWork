package leetcode19

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	fmt.Println("removeNthFromEnd3 begin")

	dummyHead := &ListNode{}
	dummyHead.Next = head
	p, q := dummyHead, dummyHead
	for i := 0; i < n+1; i++ {
		q = q.Next
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return dummyHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	fmt.Println("removeNthFromEnd2 begin")
	var temp ListNode
	temp.Next = head
	var node1 = &temp
	var node2 = &temp
	for i := 0; i < n+1; i++ {
		node1 = node1.Next
	}
	for node1 != nil {
		node1 = node1.Next
		node2 = node2.Next
	}
	node2.Next = node2.Next.Next
	return temp.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodeSlic []int = []int{head.Val}
	for head.Next != nil {
		head = head.Next
		nodeSlic = append(nodeSlic, head.Val)
	}

	var temp = ListNode{}
	var curr = &temp

	for i := 0; i < len(nodeSlic); i++ {
		if i != len(nodeSlic)-n {
			curr.Next = &ListNode{nodeSlic[i], nil}
			curr = curr.Next
		}
	}
	return temp.Next
}
