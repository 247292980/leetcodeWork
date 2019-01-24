package leetcode11

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0011(t *testing.T) {

	qs := []question{
		question{
			para{[]int{1, 2, 3, 1}},
			ans{3},
		},
		question{
			para{[]int{1, 3, 6, 4, 3, 5, 6, 7, 8, 9, 7, 5, 4, 3, 2, 1}},
			ans{48},
		},
		question{
			para{[]int{1, 1}},
			ans{1},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", maxArea(p.one))
	}
}
