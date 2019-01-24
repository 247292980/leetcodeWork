package leetcode31

func nextPermutation(nums []int) {
	var numsLen = len(nums)
	if numsLen <= 1 {
		return
	}
	//相邻两数之间，左边界为最小值就左移
	var i = numsLen - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i >= 0 {
		var j = numsLen - 1
		//找到大于nums[i]的相邻最小值
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	//交换
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	var i = start
	var j = len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
