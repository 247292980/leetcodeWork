package leetcode24

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//有个大佬和我思路一样，写的更好
func swapPairs2(head *ListNode) *ListNode {
	var cur = new(ListNode)
	var res = cur
	cur.Next = head
	for cur.Next != nil && cur.Next.Next != nil {
		t1, t2 := cur.Next, cur.Next.Next
		cur.Next, t2.Next, t1.Next, cur = t2, t1, t2.Next, t1
	}
	return res.Next
}

func swapPairs(head *ListNode) *ListNode {
	fmt.Println("SwapPairs begin")

	var cur = new(ListNode)
	var result = cur
	cur.Next = head
	for cur.Next != nil && cur.Next.Next != nil {
		/*得到要交换的两个点*/
		n1 := cur.Next
		n2 := cur.Next.Next

		cur.Next = n2

		temp := n2.Next
		/*这一步形成了一个环，debug看一下吧，挺有趣的*/
		n2.Next = n1
		n1.Next = temp

		cur = n1
	}
	return result.Next
}
