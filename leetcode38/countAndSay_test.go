package leetcode37

import (
	"fmt"
	"testing"
)

type question struct {
	para
	answer
}

// para 是参数
// one 代表第一个参数
type para struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type answer struct {
	answer string
}

func Test_Problem0038(t *testing.T) {
	qs := []question{
		question{
			para{
				3,
			},
			answer{
				"21",
			},
		},
		// 如需多个测试，可以复制上方元素。
	}
	println(1233216)

	for _, q := range qs {
		a, p := q.answer, q.para
		fmt.Printf("~~%v~~\n", p)

		fmt.Println("输入 ", p.one)
		//fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a)

		fmt.Println("结果 ", countAndSay(p.one))
	}
}
