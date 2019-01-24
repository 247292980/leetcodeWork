package leetcode32

func longestValidParentheses(s string) int {
	var left, max, temp int
	record := make([]int, len(s))

	// 统计Record,record key是）的index
	for index, value := range s {
		if value == '(' {
			left++
		} else if left > 0 {
			left--
			record[index] = 2
		}
	}

	// 修改record
	for i := 0; i < len(record); i++ {
		if record[i] == 2 {
			j := i - 1
			for record[j] != 0 {
				j--
			}
			record[i], record[j] = 1, 1
		}
	}

	// 统计结果
	for _, mapValue := range record {
		if mapValue == 0 {
			temp = 0
			continue
		}

		temp++
		if temp > max {
			max = temp
		}
	}

	return max
}
