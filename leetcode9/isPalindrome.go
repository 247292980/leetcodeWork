package leetcode9

/*第一次和官方解答一样啊。纪念下*/
func isPalindrome(x int) bool {
	var temp = x
	if temp < 0 {
		return false
	}
	var reverse = 0
	for temp/10 != 0 || temp%10 != 0 {
		reverse = reverse*10 + temp%10
		temp = temp / 10
	}
	if reverse == x {
		return true
	}
	return false
}
