package leetcode1

func twoSum(nums []int, target int) []int {
	var temp = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		needNum := target - nums[i]
		/*判断key是否存在*/
		_, ok := temp[needNum]
		if ok {
			return []int{temp[needNum], i}
		}
		temp[nums[i]] = i
	}
	return nil
}
