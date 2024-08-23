package _057_insert

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	list := [][]int{}
	for _, interval := range intervals {
		if len(list) == 0 {
			list = append(list, interval)
			continue
		}
		ia, ib := interval[0], interval[1]
		last := len(list) - 1
		nb := list[last][1]
		if ia <= nb {
			list[last][1] = max(list[last][1], ib)
			continue
		}
		list = append(list, interval)
	}
	return list
}
