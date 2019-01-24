package leetcode8

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
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
