package leetcode26

import (
	"fmt"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{1, 1, 2},
		2,
	},

	{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		5,
	},

	// 可以有多个 testcase
}

func Test_removeDuplicates(t *testing.T) {

	for _, tc := range tcs {

		fmt.Println("输入 ", tc.nums)
		fmt.Println("预期结果 ", tc.ans)
		fmt.Println("结果 ", removeDuplicates(tc.nums))
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeDuplicates(tc.nums)
		}
	}
}
