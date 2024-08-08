package _274_hIndex

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	pre := -1
	for i, citation := range citations {
		l := len(citations) - i
		if l >= citation {
			pre = citation
		} else {
			if l >= pre {
				pre = l
			}
			break
		}
	}
	return pre
}
