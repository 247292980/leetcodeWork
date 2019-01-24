package leetcode5

import (
	"fmt"
	"strings"
)

/*这个用了Manacher算法，时间复杂度最低，但是运行时间很长，这说明了算法也不是只看时间复杂度的啊*/
func longestPalindrome(s string) string {
	fmt.Println("longestPalindrome begin")
	var separatorChar = []rune("#")
	var tempSrc = separatorChar
	var rune []rune = []rune(s)
	for i := 0; i < len(rune); i++ {
		tempSrc = append(tempSrc, rune[i], separatorChar[0])
	}

	var palindrome string
	var len = len(tempSrc)
	var max = 0

	for i := 0; i < len; i++ {
		pos := 1
		for i-pos >= 0 && i+pos < len {
			if tempSrc[i-pos] != tempSrc[i+pos] {
				break
			}
			if max < pos*2+1 {
				max = pos*2 + 1
				palindrome = ""
				for j := i - pos; j <= i+pos; j++ {
					palindrome += string(tempSrc[j])
				}
			}
			pos++
		}
	}
	return strings.Replace(palindrome, "#", "", -1)
}
