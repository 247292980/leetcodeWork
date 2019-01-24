package leetcode13

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

func Test_Problem0013(t *testing.T) {

	qs := []question{
		question{
			para{"XXXIX"},
			ans{39},
		},
		question{
			para{"MDCCCLXXXVIII"},
			ans{1888},
		},
		question{
			para{"MCMLXXVI"},
			ans{1976},
		},
		question{
			para{"MMMCMXCIX"},
			ans{3999},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", romanToInt(p.one))
	}
}
