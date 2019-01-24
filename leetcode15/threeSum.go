package leetcode15

import (
	"sort"
)

/*开头的时候思路是对的，但是双指针+定点的定点找错了...后模仿js重写一次*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var threeSum [][]int

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			l := i + 1
			r := len(nums) - 1
			for l < r {
				s := nums[i] + nums[l] + nums[r]
				if s == 0 {
					threeSum = append(threeSum, []int{nums[l], nums[i], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for r > l && nums[r] == nums[r+1] {
						r--
					}
				} else if s > 0 {
					r--
				} else {
					l++
				}
			}
		}
	}
	return threeSum
}
