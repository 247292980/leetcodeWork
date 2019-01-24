package leetcode16

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

func Test_Problem0016(t *testing.T) {

	qs := []question{

		question{
			para{[]int{-1, 2, 1, -4}, 1},
			ans{2},
		},
		question{
			para{[]int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4}, 1},
			ans{0},
		},
		question{
			para{[]int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4}, 0},
			ans{0},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", threeSumClosest(p.one, p.two))
	}
}
