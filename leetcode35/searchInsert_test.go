package leetcode35

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
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0035(t *testing.T) {

	qs := []question{

		question{
			para{[]int{1, 3, 5, 6}, 5},
			ans{2},
		},

		question{
			para{[]int{1, 3, 5, 6}, 2},
			ans{1},
		},

		question{
			para{[]int{1, 3, 5, 6}, 7},
			ans{4},
		},

		question{
			para{[]int{1, 3, 5, 6}, 0},
			ans{0},
		},

		question{
			para{[]int{1, 3, 5, 6}, 6},
			ans{3},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a)
		fmt.Println("结果 ", searchInsert(p.one, p.two))
	}
}
