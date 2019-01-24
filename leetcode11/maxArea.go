package leetcode11

func maxArea(height []int) int {
	var r = len(height) - 1
	var l = 0
	var temp = 0
	var maxArea = 0

	for l < r {
		if height[l] < height[r] {
			temp = height[l] * (r - l)
			l++
		} else {
			temp = height[r] * (r - l)
			r--
		}
		if temp > maxArea {
			maxArea = temp
		}
	}
	return maxArea
}
