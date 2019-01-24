package leetcode3

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	var temp = make(map[rune]int)
	var rune []rune = []rune(s)
	var len = len(rune)
	var max float64 = 0

	for i, j := float64(0), 0; j < len; j++ {
		_, ok := temp[rune[j]]
		if ok {
			/*记录当前字符上一次出现的位置*/
			i = math.Max(float64(temp[rune[j]]), i)
		}
		max = math.Max(max, float64(j)+1-i)
		/*刷新记录当前字符的位置*/
		temp[rune[j]] = j + 1
	}
	return int(max)
}
