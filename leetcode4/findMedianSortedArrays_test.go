package leetcode4

import (
	"fmt"
	"testing"
)

type para struct {
	one []int
	two []int
}

type ans struct {
	one float64
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {

	qs := []question{
		question{
			p: para{
				one: []int{1, 3},
				two: []int{2},
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{1, 3},
				two: []int{2, 4},
			},
			a: ans{
				one: 2.5,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		fmt.Println("输入 ", p.one)
		fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", findMedianSortedArrays(p.one, p.two))
	}

}
