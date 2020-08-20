package problem0062

func uniquePaths(m int, n int) int {
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
		if i == 0 {
			for j := 0; j < n; j++ {
				paths[i][j] = 1
			}
		} else {
			paths[i][0] = 1
			for j := 1; j < n; j++ {
				paths[i][j] = paths[i][j-1] + paths[i-1][j]
			}
		}
	}
	return paths[m-1][n-1]
}
