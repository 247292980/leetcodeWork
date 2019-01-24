package leetcode29

import (
	"fmt"
	"math"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0029(t *testing.T) {

	qs := []question{

		question{
			para{2, 3},
			ans{0},
		},

		question{
			para{2, 0},
			ans{math.MaxInt32},
		},

		question{
			para{1024, 3},
			ans{341},
		},

		question{
			para{1024, -3},
			ans{-341},
		},

		question{
			para{-1024, 3},
			ans{-341},
		},

		question{
			para{-1024, -3},
			ans{341},
		},

		question{
			para{1024, 1},
			ans{1024},
		},

		question{
			para{2147483647, 1},
			ans{2147483647},
		},

		question{
			para{2147483648, 1},
			ans{math.MaxInt32},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a)
		fmt.Println("结果 ", divide(p.one, p.two))
	}
}
