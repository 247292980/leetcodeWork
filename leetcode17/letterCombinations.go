package leetcode17

import "fmt"

func letterCombinations(digits string) []string {
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
