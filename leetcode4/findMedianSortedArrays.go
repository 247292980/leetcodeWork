package leetcode4

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	fmt.Println("findMedianSortedArrays begin")
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	var midPoint = len(nums) / 2

	if len(nums)%2 == 0 {
		return float64(nums[midPoint]+nums[midPoint-1]) / 2
	} else {
		return float64(nums[midPoint])
	}
}
