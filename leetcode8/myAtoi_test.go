package leetcode8

import (
	"fmt"
	"math"
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
				one: "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459",
			},
			a: ans{
				one: math.MaxInt32,
			},
		},

		question{
			p: para{
				one: "123",
			},
			a: ans{
				one: 123,
			},
		},
		question{
			p: para{
				one: "-123",
			},
			a: ans{
				one: -123,
			},
		},
		question{
			p: para{
				one: "2147483648",
			},
			a: ans{
				one: math.MaxInt32,
			},
		},
		question{
			p: para{
				one: "-2147483649",
			},
			a: ans{
				one: math.MinInt32,
			},
		},
		question{
			p: para{
				one: " 1234a6789",
			},
			a: ans{
				one: 1234,
			},
		},
		question{
			p: para{
				one: "  -0012a42 ",
			},
			a: ans{
				one: -12,
			},
		},
		question{
			p: para{
				one: " asdfdfs ",
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: " ",
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: " +1   ",
			},
			a: ans{
				one: 1,
			},
		},

		question{
			p: para{
				one: "-",
			},
			a: ans{
				one: 0,
			},
		},

		question{
			p: para{
				one: "922337999995452345782348957234895793875923845789234758923745987239485798345789237598235980234859023849058349058903890869059068490683490869038690385690385906839056890548690586904568905468905908590839056890345869034856903568903854690835906834906839045869034869034568903458690356903569056908345906839056890586903546890345869034568903586905689054685690905689035468905689056890879056907890879086903548690387905469054690890689035869038569034856908356908345906890345869056890356890358690569083546908549086905690345869038569034569083590689058690385690358690586908345906839086390689056903869058690345869038690586908569054690834590689054869083569035490689035689058690586905409689086905869038569083549068390468903586903569038549068905869054690345906890346904856903546908345906890568903569054690590685906905689058690586905869056890869035869035890789068790907903890835657428975457575789075098759084752897589475029847589047589234759028475902847592834752908759827589725987517891598715908749871908579841790817598715901875901874190879085791879018571897491837249874987235987589734897123489712390847913857190287549018735902036854775809",
			},
			a: ans{
				one: math.MaxInt32,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		fmt.Println("输入 ", p.one)
		fmt.Println("预期结果 ", a.one)
		fmt.Println("结果 ", myAtoi(p.one))
	}
}
