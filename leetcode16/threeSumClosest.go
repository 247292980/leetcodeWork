package leetcode16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var ThreeSumClosest = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			l := i + 1
			r := len(nums) - 1
			for l < r {
				if ThreeSumClosest == target {
					return target
				}
				s := nums[i] + nums[l] + nums[r]
				if math.Abs(float64(s-target)) < math.Abs(float64(ThreeSumClosest-target)) {
					ThreeSumClosest = s
				}
				if s > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return ThreeSumClosest
}
