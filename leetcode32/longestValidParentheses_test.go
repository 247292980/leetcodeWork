package leetcode32

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
	one string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0032(t *testing.T) {

	qs := []question{
		question{
			para{")(()()))((((())))("},
			ans{8},
		},

		question{
			para{"()(()"},
			ans{2},
		},

		question{
			para{"(()())"},
			ans{6},
		},

		question{
			para{"(()"},
			ans{2},
		},

		question{
			para{")()())"},
			ans{4},
		},

		question{
			para{"((()))"},
			ans{6},
		},

		question{
			para{"((())))(((())))"},
			ans{8},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		//fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a)
		fmt.Println("结果 ", longestValidParentheses(p.one))
	}
}
