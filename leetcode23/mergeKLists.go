package leetcode23

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 很难看了快一个小时了
**/
func mergeKLists2(lists []*ListNode) *ListNode {
	fmt.Println("mergeKLists2 begin")
	/*end first是同一个指针*/
	var first, end *ListNode
	size := len(lists)

	for {
		index := -1
		var nums []int
		/*找出lists最小的值*/
		for i := 0; i < size; i++ {
			if lists[i] != nil {
				if index == -1 || lists[index].Val > lists[i].Val {
					index = i
					nums = nums[:0]
				} else if lists[index].Val == lists[i].Val {
					nums = append(nums, i)
				}
			}
		}
		if index == -1 {
			break
		} else {
			if first == nil {
				first = lists[index]
				end = lists[index]
			} else {
				end.Next = lists[index]
				end = lists[index]
			}
			lists[index] = lists[index].Next
			size := len(nums)
			if size > 0 {
				for i := 0; i < size; i++ {
					/*end.Next = lists[nums[i]] 改了end和first，把最小的数放到first里面*/
					end.Next = lists[nums[i]]
					/*end.=lists[nums[i]] 只改了end*/
					end = lists[nums[i]]
					/*已经放到first的指针移位*/
					lists[nums[i]] = lists[nums[i]].Next
				}
			}
		}
	}
	return first
}
func mergeKLists(lists []*ListNode) *ListNode {
	fmt.Println("mergeKLists begin")
	if lists == nil {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var result = mergeTwoLists(lists[0], lists[1])
	for i := 2; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}
	return result
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
