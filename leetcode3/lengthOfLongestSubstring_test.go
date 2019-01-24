package leetcode3

import (
	"fmt"
	"testing"
)

type para struct {
	one string
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
				one: "abcabcbb",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "bbbbbbbb",
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: "pwwkew",
			},
			a: ans{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		fmt.Println("输入 ", p.one)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", lengthOfLongestSubstring(p.one))
	}
}
