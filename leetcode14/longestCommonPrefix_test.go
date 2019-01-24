package leetcode14

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
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one string
}

func Test_Problem0014(t *testing.T) {

	qs := []question{
		question{
			para{
				[]string{"abcdd", "abcde", "ab"},
			},
			ans{"ab"},
		},
		question{
			para{
				[]string{"abcdd", "abcde"},
			},
			ans{"abcd"},
		},
		question{
			para{
				[]string{},
			},
			ans{""},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", longestCommonPrefix(p.one))
	}
}
