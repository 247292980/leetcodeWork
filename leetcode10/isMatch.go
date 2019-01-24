package leetcode10

import "fmt"

func isMatch(s string, p string) bool {
	fmt.Println("IsMatch begin")

	sLen := len(s)
	pLen := len(p)
	dp := make([][]bool, sLen+1, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([]bool, pLen+1, pLen+1)
		if i != 0 {
			dp[i][0] = false // s非空，p为空时肯定不匹配
		}
	}
	dp[0][0] = true // 都为空时认为是匹配
	for i := 0; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] == '*' {
				// 即*之前的字符重复0次后匹配（等价于dp[i][j-2]）或者重复一次后匹配（等价于dp[i-1][j]）
				dp[i][j] = dp[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j])
			} else {
				// 非*结尾时，对比最后一位是否可匹配，如果匹配则当前dp[i][j]等价于dp[i-1][j-1]
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return dp[sLen][pLen]
}
