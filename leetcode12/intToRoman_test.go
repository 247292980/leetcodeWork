package leetcode12

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
	one int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one string
}

func Test_Problem0012(t *testing.T) {

	qs := []question{
		question{
			para{39},
			ans{"XXXIX"},
		},
		question{
			para{1888},
			ans{"MDCCCLXXXVIII"},
		},
		question{
			para{1976},
			ans{"MCMLXXVI"},
		},
		question{
			para{3999},
			ans{"MMMCMXCIX"},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		fmt.Println("输入 ", p.one)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", intToRoman(p.one))
	}
}
