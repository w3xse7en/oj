package problem0583

func minDistance(word1 string, word2 string) int {
	m, n := len(word1)+1, len(word2)+1
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if y == 0 && x != 0 {
				dp[y][x] = dp[y][x-1] + 1
			}
			if x == 0 && y != 0 {
				dp[y][x] = dp[y-1][x] + 1
			}
			if y >= 1 && x >= 1 {
				t := min(dp[y][x-1], dp[y-1][x]) + 1
				if word1[y-1] == word2[x-1] {
					t = dp[y-1][x-1]
				}
				dp[y][x] = t
			}
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
