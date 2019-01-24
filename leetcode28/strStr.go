package leetcode28

import "strings"

/*kmp算法,尴尬的是和index速度一样...*/
func strStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	var hp = 0
	var np = 0
	for hp < len(haystack) {
		if haystack[hp] == needle[np] {
			np++
		} else {
			hp = hp - np
			np = 0
		}
		hp++
		if np == len(needle) {
			return hp - len(needle)
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	return strings.Index(haystack, needle)
}
