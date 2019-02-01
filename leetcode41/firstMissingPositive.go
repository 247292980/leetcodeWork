package leetcode37

import "sort"

/*
一开始简单难度用下set就好了
然后看到哦 只能使用常数级别的空间
不愧是困难
*/
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	index := 0
	compareIndex := 0

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			/*第一个大于0的数的index*/
			index = next(nums, i)
			compareIndex = next(nums, index+1)
			if nums[index] > 1 {
				return 1
			}
			break
		}
	}
	/*全部数字为负*/
	if nums[index] < 1 {
		return 1
	}
	/*找中间一个不是递增1*/
	for compareIndex < len(nums) && nums[index]+1 == nums[compareIndex] {
		index = compareIndex
		compareIndex = next(nums, index+1)
		/*都是递增的*/
		if index+1 == len(nums) {
			return nums[index] + 1
		}
	}
	return nums[index] + 1
}

/*跳过重复*/
func next(nums []int, index int) int {
	if index+1 >= len(nums) {
		return index
	}

	for index+1 < len(nums) && nums[index] == nums[index+1] {
		index++
	}
	return index
}
