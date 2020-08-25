package problem0077

func combine(n int, k int) [][]int {
	var result [][]int
	for i := 1; i <= n; i++ {
		dsf(i, n, 0, k, []int{}, &result)
	}
	return result
}
func dsf(i, n, step, end int, tmp []int, result *[][]int) {
	if step != end {
		tmp = append(tmp, i)
		if len(tmp) == end {
			t := make([]int, len(tmp))
			copy(t, tmp)
			*result = append(*result, t)
		}
		for j := i + 1; j <= n; j++ {
			dsf(j, n, step+1, end, tmp, result)
		}
	}
}
