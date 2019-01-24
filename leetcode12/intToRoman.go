package leetcode12

import (
	"bytes"
	"fmt"
)

func intToRoman(num int) string {
	fmt.Println("intToRoman begin")

	bufStr := bytes.NewBufferString("")
	nums := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNums := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < 13; i++ {
		tmp := num / nums[i]
		if tmp > 0 {
			bufStr.Write(bytes.Repeat([]byte(romanNums[i]), tmp))
			num -= tmp * nums[i]
		} else {
			continue
		}
	}
	return bufStr.String()
}
