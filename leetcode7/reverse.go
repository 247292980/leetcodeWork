package leetcode7

func reverse(x int) int {
	var reverse = 0
	for x/10 != 0 || x%10 != 0 {
		reverse = reverse*10 + x%10
		x = x / 10
	}
	if reverse > (2<<30)-1 || reverse < -2<<30 {
		reverse = 0
	}
	return reverse
}
