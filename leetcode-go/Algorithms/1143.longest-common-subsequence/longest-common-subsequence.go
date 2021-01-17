package problem1143

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1)+1, len(text2)+1
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			t := max(dp[i-1][j], dp[i][j-1])
			if text1[i-1] == text2[j-1] {
				t = dp[i-1][j-1] + 1
			}
			dp[i][j] = t
		}
	}
	return dp[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
