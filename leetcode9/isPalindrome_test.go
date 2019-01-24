package leetcode9

import (
	"fmt"
	"testing"
)

type question struct {
	p para
	a ans
}

type para struct {
	one int
}

type ans struct {
	one bool
}

func Test_Problem0009(t *testing.T) {

	qs := []question{
		question{
			p: para{
				one: 12321,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: 1231,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: -12321,
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		fmt.Println("输入 ", p.one)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", isPalindrome(p.one))
	}
}
