package leetcode37

import "sort"

/*暴力法,这题难度比上题还高？有没有搞错，，，
能用暴力法还100%的题其实没有难度desu*/
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	solution := []int{}

	find(candidates, solution, target, &res)
	return res
}

func find(candidates []int, solution []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, solution)
	}
	if len(candidates) == 0 || target < candidates[0] {
		return
	}
	solution = solution[:len(solution):len(solution)]

	find(candidates, append(solution, candidates[0]), target-candidates[0], res)

	find(candidates[1:], solution, target, res)
}
