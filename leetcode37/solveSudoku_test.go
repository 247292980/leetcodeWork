package leetcode37

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
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one [][]byte
}

func Test_Problem0037(t *testing.T) {

	qs := []question{

		question{
			para{[][]byte{
				[]byte("..9748..."),
				[]byte("7........"),
				[]byte(".2.1.9..."),
				[]byte("..7...24."),
				[]byte(".64.1.59."),
				[]byte(".98...3.."),
				[]byte("...8.3.2."),
				[]byte("........6"),
				[]byte("...2759.."),
			}},
			ans{[][]byte{
				[]byte("519748632"),
				[]byte("783652419"),
				[]byte("426139875"),
				[]byte("357986241"),
				[]byte("264317598"),
				[]byte("198524367"),
				[]byte("975863124"),
				[]byte("832491756"),
				[]byte("641275983"),
			}},
		},
		// 如需多个测试，可以复制上方元素。
	}
	println(1233216)

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		fmt.Println("输入 ", p.one)
		//fmt.Println("输入 ", p.two)
		fmt.Println("预期结果 ", a)
		solveSudoku(p.one)
		fmt.Println("结果 ", p.one)
	}
}
