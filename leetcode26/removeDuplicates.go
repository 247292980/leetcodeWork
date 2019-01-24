package leetcode26

func removeDuplicates2(nums []int) int {

	length := len(nums)
	if length == 0 {
		return 0
	}

	i := 0
	for j := 1; j < length; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
func removeDuplicates(nums []int) int {

	length := len(nums)
	i := 0
	for i < length-1 {
		if nums[i] == nums[i+1] {
			nums = append(nums[0:i+1], nums[i+2:]...)
			length--
		} else {
			i++
		}
	}
	return length
}
