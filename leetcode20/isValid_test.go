package leetcode20

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
	one bool
}

func Test_Problem0020(t *testing.T) {

	qs := []question{

		question{
			para{"()[]{}"},
			ans{true},
		},
		question{
			para{"(]"},
			ans{false},
		},
		question{
			para{"({[]})"},
			ans{true},
		},
		question{
			para{"(){[({[]})]}"},
			ans{true},
		},
		question{
			para{"((([[[{{{"},
			ans{false},
		},
		question{
			para{"(())]]"},
			ans{false},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		//fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", isValid(p.one))
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 1; i < b.N; i++ {
		isValid("{{{{{[[[[[((((()))))]]]]]}}}}}")
	}
}
