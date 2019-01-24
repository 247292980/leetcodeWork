package leetcode35

func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}

	for index, value := range nums {
		if value >= target {
			return index
		}
	}
	return len(nums)
}
