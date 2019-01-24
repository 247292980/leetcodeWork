package leetcode27

/*开窍了，造成之前60%70%的主要原因是没有把一些特殊情况列出*/
func removeElement(nums []int, val int) int {

	length := len(nums)
	if length <= 0 {
		return 0
	}
	for i := 0; i < length; i++ {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
			i--
			length--
		}
	}
	return length
}
