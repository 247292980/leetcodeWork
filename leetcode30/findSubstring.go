package leetcode30

import (
	"fmt"
	"strings"
)

/*超时的写出来还以为很简单，然后看这个正确答案...*/
func findSubstring2(s string, words []string) []int {
	sLen := len(s)
	result := make([]int, 0, sLen)

	wordsLen := len(words)
	if sLen == 0 || wordsLen == 0 || sLen < wordsLen*len(words[0]) {
		return result
	}

	interval := len(words[0])

	// record 记录 words 中每个单词总共出现的次数，该数据并不会变动
	record := make(map[string]int, wordsLen)
	for _, w := range words {
		if len(w) != interval {
			// 新添加的 example 2 出现了 words 中单词长度不一致的情况。
			// 这个违反了假设
			// 直接返回 result
			return result
		}
		record[w]++
	}

	// remain 记录 words 中每个单词还能出现的次数
	remain := make(map[string]int, wordsLen)
	// count 记录符合要求的单词的连续出现次数
	count := 1 // count 的初始值只要不为 0 就可以
	left, right := 0, 0

	/**
	 * s[left:right] 作为一个窗口存在
	 * 假设 word:= s[right:right+interval]
	 * 如果 remain[word]>0 ，那么移动窗口 右边，同时更新移动后 s[left:right] 的统计信息
	 * 		remain[word]--
	 * 		right += interval
	 * 		count++
	 * 否则，移动窗口 左边，同时更新移动后 s[left:right] 的统计信息
	 * 		remain[s[left:left+interval]]++
	 * 		count--
	 * 		left += interval
	 *
	 * 每次移动窗口 右边 后，如果 count = wordsLen ，那么
	 * 		说明找到了一个符合条件的解
	 * 		append(result, left)，然后
	 * 		移动窗口 左边
	 */

	// reset 重置 remain 和 count
	reset := func() {
		if count == 0 {
			// 统计记录没有被修改，不需要重置
			// 因为有这个预判，所以需要第一次执行 reset 时
			// count 的值不能为 0
			// 即 count 的初始值不能为 0
			return
		}
		for k, v := range record {
			remain[k] = v
		}
		count = 0
	}

	// moveLeft 让 left 指向下一个单词
	//大佬思维可能和我想的不一样moveLeft我第一眼就觉得的左移，其实是左指针右移
	moveLeft := func() {
		// left 指向下一个单词前，需要修改统计记录
		// 增加 left 指向的 word 可出现次数一次，
		remain[s[left:left+interval]]++
		// 连续符合条件的单词数减少一个
		count--
		// left 后移一个 word 的长度
		left += interval
	}

	// left 需要分别从{0,1,...,interval-1}这些位置开始检查，才能不遗漏
	for i := 0; i < interval; i++ {
		left, right = i, i
		reset()

		// s[left:] 的长度 >= words 中所有 word 组成的字符串的长度时，
		// s[left:] 中才有可能存在要找的字符串
		for sLen-left >= wordsLen*interval {
			word := s[right : right+interval]
			remainTimes, ok := remain[word]

			switch {
			case !ok:
				// word 不在 words 中
				// 从 right+interval 处，作为新窗口，重新开始统计
				left, right = right+interval, right+interval
				reset()
			case remainTimes == 0:
				// word 的出现次数上次就用完了
				// 说明 word 在 s[left:right] 中出现次数过多
				moveLeft()
				// 这个case会连续出现
				// 直到 s[left:right] 中的统计结果是 remain[word] == 1
				// 这个过程中，right 一直不变
			default:
				// ok && remainTimes > 0，word 符合出现的条件
				// moveRight
				remain[word]--
				count++
				right += interval
				// 检查 words 能否排列组合成 s[left:right]
				if count == wordsLen {
					result = append(result, left)
					// moveLeft 可以避免重复统计 s[left+interval:right] 中的信息
					moveLeft()
				}
			}
		}
	}

	return result
}

/*超时超时超时超时*/
func findSubstring(s string, words []string) []int {
	fmt.Println("FindSubstring begin")

	var result []int

	if len(s) == 0 || len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return result
	}
	for i := 0; i < len(s); i++ {
		find(&result, s[i:], words, i)
	}
	return result
}

func find(result *[]int, s string, words []string, cur int) {
	for index, value := range words {
		i := strings.Index(s, value)
		/*第一个就是words里的*/
		if i == 0 {
			var tempS = s[len(value):]
			var tempWords []string
			tempWords = append(tempWords, words[:]...)
			tempWords = append(tempWords[:index], tempWords[index+1:]...)
			if len(tempWords) == 0 {
				var ok = true
				//去重
				for _, value := range *result {
					if !ok {
						break
					}
					if value == cur {
						ok = false
					}
				}

				if ok {
					*result = append(*result, cur)
				}
				return
			}

			find(result, tempS, tempWords, cur)
		}
	}
}
