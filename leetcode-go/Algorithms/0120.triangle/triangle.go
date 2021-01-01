package problem0120

func minimumTotal(triangle [][]int) int {
	for lv, ints := range triangle {
		if lv == 0 {
			continue
		}
		for idx, v := range ints {
			pre := triangle[lv-1]
			i, j := idx-1, idx
			best := 0xffffffff
			if i >= 0 && i < len(pre) && best > v+pre[i] {
				best = v + pre[i]
			}
			if j >= 0 && j < len(pre) && best > v+pre[j] {
				best = v + pre[j]
			}
			ints[idx] = best
		}
	}
	min := 0xffffffff
	for _, v := range triangle[len(triangle)-1] {
		if min > v {
			min = v
		}
	}
	return min
}
