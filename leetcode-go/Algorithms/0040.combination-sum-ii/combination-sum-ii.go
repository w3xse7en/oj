package problem0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	combinations := [][]int{}
	dfs(target, candidates, []int{}, &combinations)
	return combinations
}

func dfs(target int, candidates, combination []int, combinations *[][]int) {
	for i := range candidates {
		if i != 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if target >= candidates[i] {
			combination := append(combination, candidates[i])
			if target == candidates[i] {
				t := make([]int, len(combination))
				copy(t, combination)
				*combinations = append(*combinations, t)
				return
			}
			dfs(target-candidates[i], candidates[i+1:], combination, combinations)
		} else {
			return
		}
	}
}
