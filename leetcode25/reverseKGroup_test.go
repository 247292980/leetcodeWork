package leetcode25

import (
	"fmt"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	k    int
	ans  []int
}{

	{
		[]int{1, 2, 3, 4, 5},
		3,
		[]int{3, 2, 1, 4, 5},
	},

	{
		[]int{1, 2, 3, 4, 5},
		1,
		[]int{1, 2, 3, 4, 5},
	},

	// 可以有多个 testcase
}

func Test_reverseKGroup(t *testing.T) {
	for _, tc := range tcs {
		head := s2l(tc.head)
		ans := s2l(tc.ans)

		fmt.Println("输入 ", head)
		fmt.Println("输入 ", tc.k)
		fmt.Println("预期结果 ", ans)
		fmt.Println("结果 ", reverseKGroup(head, tc.k))
	}
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

func Benchmark_reverseKGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := s2l(tc.head)
			reverseKGroup(head, tc.k)
		}
	}
}
