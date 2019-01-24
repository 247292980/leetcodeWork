package leetcode7

import (
	"fmt"
	"testing"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {

	qs := []question{
		question{
			p: para{
				one: 123,
			},
			a: ans{
				one: 321,
			},
		},
		question{
			p: para{
				one: -123,
			},
			a: ans{
				one: -321,
			},
		},
		question{
			p: para{
				one: 1234567899,
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: -1234567899,
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		fmt.Println("输入 ", p.one)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", reverse(p.one))
	}
}
