package leetcode22

func generateParenthesis(n int) []string {
	var result []string
	generate(&result, "", 0, 0, n)
	return result
}
func generate(result *[]string, s string, left, right, n int) {
	if left > n || right > n {
		return
	}
	if right == n && left == n {
		*result = append(*result, s)
		return
	}

	if left >= right {
		generate(result, s+"(", left+1, right, n)
		generate(result, s+")", left, right+1, n)
	}
}
