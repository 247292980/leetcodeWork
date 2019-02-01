package leetcode37

import "sort"

/*上一题珠玉在前啊，突然发现解法都固定了啊*/
func combinationSum2(candidates []int, target int) [][]int {
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
	find(candidates[1:], append(solution, candidates[0]), target-candidates[0], res)
	find(next(candidates), solution, target, res)
}

func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}
