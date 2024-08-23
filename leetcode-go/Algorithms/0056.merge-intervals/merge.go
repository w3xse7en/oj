package problem0056

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	list := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		a, b := intervals[i][0], intervals[i][1]
		last := len(list) - 1
		lb := list[last][1]
		if a <= lb {
			list[last][1] = max(list[last][1], b)
			continue
		}
		list = append(list, intervals[i])
	}
	return list
}
