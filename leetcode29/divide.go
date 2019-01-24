package leetcode29

import "math"

/*震惊！我看到res - res - res就知道大神之所以是大神*/
func divide2(m, n int) int {
	// 防止有人把0当做除数
	if n == 0 {
		return math.MaxInt32
	}

	signM, absM := analysis(m)
	signN, absN := analysis(n)

	res, _ := d(absM, absN, 1)

	// 修改res的符号
	if signM != signN {
		res = res - res - res
	}

	// 检查溢出
	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
func analysis(num int) (sign, abs int) {
	sign = 1
	abs = num
	if num < 0 {
		sign = -1
		abs = num - num - num
	}
	return
}

// d 计算m/n的值，返回结果和余数
// m >= 0 分子
// n > 0 分母
// count == 1, 代表初始n的个数，在递归过程中，count == 当前的n/初始的n
func d(m, n, count int) (res, remainder int) {
	switch {
	case m < n:
		return 0, m
	case m <= n && m < n+n:
		return count, m - n
	default:
		res, remainder = d(m, n+n, count+count)
		if remainder >= n {
			return res + count, remainder - n
		}
		return
	}
}

/*一次过的都不写测试用例了...
orz,题目说不能用除法...*/
func divide(dividend int, divisor int) int {
	var result = dividend / divisor
	if divisor == 0 || dividend == 0 {
		return 0
	}
	// 检查溢出
	if result < math.MinInt32 || result > math.MaxInt32 {
		return math.MaxInt32
	}
	return result
}
