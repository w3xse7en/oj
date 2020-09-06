package problem0039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	dfs(target, candidates, []int{}, &result)
	return result
}
func dfs(target int, candidates, combination []int, result *[][]int) {
	for i := range candidates {
		if target >= candidates[i] {
			if target == candidates[i] {
				t := make([]int, len(combination)+1)
				copy(t, append(combination, candidates[i]))
				*result = append(*result, t)
				return
			}
			dfs(target-candidates[i], candidates[i:], append(combination, candidates[i]), result)
		} else {
			return
		}
	}
}
