package leetcode27

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
	one []int
}

func Test_Problem0027(t *testing.T) {

	qs := []question{

		question{
			para{[]int{3, 2, 2, 3}, 3},
			ans{[]int{2, 2}},
		},

		question{
			para{[]int{3, 1, 5, 7, 2, 2, 3}, 3},
			ans{[]int{2, 1, 5, 7, 2}},
		},

		question{
			para{[]int{1, 5, 7, 2, 2}, 3},
			ans{[]int{1, 5, 7, 2, 2}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a)
		fmt.Println("结果 ", removeElement(p.one, p.two))
	}
}
