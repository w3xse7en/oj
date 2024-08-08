package _207_canFinish

import (
	"slices"
)

var okPath map[int]bool

func canFinish(numCourses int, prerequisites [][]int) bool {
	mp := map[int][]int{}
	okPath = map[int]bool{}
	for _, prerequisite := range prerequisites {
		a, b := prerequisite[0], prerequisite[1]
		mp[a] = append(mp[a], b)
	}

	for k := range mp {
		if dfs(k, []int{}, mp) {
			return false
		}
	}
	return true
}

func dfs(ori int, visit []int, mp map[int][]int) bool {
	if okPath[ori] {
		return false
	}
	okPath[ori] = true
	for _, next := range mp[ori] {
		if slices.Index(visit, next) != -1 {
			return true
		}
		if dfs(next, append(visit, next), mp) {
			return true
		}
	}
	return false
}
