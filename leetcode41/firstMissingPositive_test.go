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
	one []int
	//two int
}

// ans 是答案
// one 代表第一个答案
type answer struct {
	answer int
}

func Test_Problem0041(t *testing.T) {
	qs := []question{
		question{
			para{
				one: []int{1, 2, 0},
			},
			answer{
				3,
			},
		},
		question{
			para{
				one: []int{},
			},
			answer{
				1,
			},
		}, question{
			para{
				one: []int{1},
			},
			answer{
				2,
			},
		}, question{
			para{
				one: []int{3, 4, -1, 1},
			},
			answer{
				2,
			},
		}, question{
			para{
				one: []int{7, 8, 9, 11, 12},
			},
			answer{
				1,
			},
		}, question{
			para{
				one: []int{1, 1000},
			},
			answer{
				2,
			},
		}, question{
			para{
				one: []int{-1, 4, 2, 1, 9, 10},
			},
			answer{
				3,
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

		fmt.Println("结果 ", firstMissingPositive(p.one))
	}
}
