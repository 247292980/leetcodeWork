package leetcode10

import (
	"fmt"
	"testing"
)

type question struct {
	p para
	a ans
}

type para struct {
	one string
	two string
}

type ans struct {
	one bool
}

func Test_Problem0010(t *testing.T) {

	qs := []question{

		question{
			p: para{
				one: "aa",
				two: "a",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "aa",
				two: "aa",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aaa",
				two: "aa",
			},
			a: ans{
				one: false,
			},
		},

		question{
			p: para{
				one: "aa",
				two: "a*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aa",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aab",
				two: "c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aaaaaaaab",
				two: "c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*c",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "ab",
				two: "z*t*x*c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: "c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "abc",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*b.*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "b",
				two: ".*b.*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "b",
				two: ".*...b",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "b",
				two: ".*..*b",
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		fmt.Println("输入 ", p.one)
		fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", isMatch(p.one, p.two))
	}
}
