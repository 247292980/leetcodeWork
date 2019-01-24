package leetcode5

import (
	"fmt"
	"testing"
)

type para struct {
	one string
}

type ans struct {
	one string
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {

	qs := []question{
		question{
			p: para{
				one: "babad",
			},
			a: ans{
				one: "bab",
			},
		},
		question{
			p: para{
				one: "cbbd",
			},
			a: ans{
				one: "bb",
			},
		},
		question{
			p: para{
				one: "abbcccddcccbba",
			},
			a: ans{
				one: "abbcccddcccbba",
			},
		},
		question{
			p: para{
				one: "a",
			},
			a: ans{
				one: "a",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		fmt.Println("输入 ", p.one)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", longestPalindrome(p.one))
	}
}
