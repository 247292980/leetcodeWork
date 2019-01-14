package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func LetterCombinations(digits string) []string {
	fmt.Println("letterCombinations begin")
	if len(digits) == 0 {
		return nil
	}
	var LetterCombinations []string
	phoneMap := make(map[byte][]byte)
	phoneMap['2'] = []byte{'a', 'b', 'c'}
	phoneMap['3'] = []byte{'d', 'e', 'f'}
	phoneMap['4'] = []byte{'g', 'h', 'i'}
	phoneMap['5'] = []byte{'j', 'k', 'l'}
	phoneMap['6'] = []byte{'m', 'n', 'o'}
	phoneMap['7'] = []byte{'p', 'q', 'r', 's'}
	phoneMap['8'] = []byte{'t', 'u', 'v'}
	phoneMap['9'] = []byte{'w', 'x', 'y', 'z'}

	letterCombination(digits, phoneMap, 0, "", &LetterCombinations)

	return LetterCombinations
}

func letterCombination(digits string, phoneMap map[byte][]byte, level int, resItem string, res *[]string) {
	if level == len(digits) {
		*res = append(*res, resItem)
		return
	}
	itemBytes := phoneMap[digits[level]]
	for i := 0; i < len(itemBytes); i++ {
		letterCombination(digits, phoneMap, level+1, resItem+string(itemBytes[i]), res)
	}
}
func LongestCommonPrefix(strs []string) string {
	fmt.Println("LongestCommonPrefix begin")

	if len(strs) == 0 {
		return ""
	}
	short := strs[0]
	for _, s := range strs {
		if len(short) > len(s) {
			short = s
		}
	}

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}
	return short
}

func IsMatch(s string, p string) bool {
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

func MyAtoi(str string) int {
	fmt.Println("MyAtoi begin")
	arr := "+-1234567890"
	str = strings.TrimSpace(str)
	length := len(str)
	if length == 0 {
		return 0
	}

	if !strings.Contains(arr, str[0:1]) {
		return 0
	}

	result, err := strconv.Atoi(str)
	if err == nil {
		return checkInt32(result)
	}

	isNev := str[0:1] == "-" || str[0:1] == "+"
	for i := 1; i <= length; i++ {
		if isNev && i == 1 {
			continue
		}
		isNumber, result := isNumber(str, 0, i)
		if !isNumber {
			return result
		}
		result = checkInt32(result)
	}
	return 0
}

func isNumber(str string, start, end int) (bool, int) {
	s := str[start:end]
	i, err := strconv.Atoi(s)
	if err != nil {
		return false, 0
	}
	return true, i
}

func checkInt32(r int) int {
	if r > math.MaxInt32 {
		return math.MaxInt32
	}
	if r < math.MinInt32 {
		return math.MinInt32
	}
	return r
}
func ConvertZ(s string, numRows int) string {
	fmt.Println("convertZ begin")
	if numRows == 1 {
		return s
	}

	var rune []rune = []rune(s)
	var zStr string
	var separatorLen = numRows*2 - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += separatorLen {
			zStr = zStr + string(rune[j+i])
			if i != 0 && i != numRows-1 && j+separatorLen-i < len(s) {
				zStr = zStr + string(rune[j+separatorLen-i])
			}
		}
	}
	return zStr
}

/*这个用了Manacher算法，时间复杂度最低，但是运行时间很长，这说明了算法也不是只看时间复杂度的啊*/
func LongestPalindrome(s string) string {
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
func LengthOfLongestSubstring(s string) int {
	fmt.Println("lengthOfLongestSubstring begin")
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
