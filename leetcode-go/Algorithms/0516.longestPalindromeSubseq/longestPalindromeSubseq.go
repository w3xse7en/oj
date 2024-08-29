package _516_longestPalindromeSubseq

func longestPalindromeSubseq(s string) int {
	l := len(s)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
	}
	for path := 0; path < l; path++ {
		for i, j := 0, path; i < l && j < l; i++ {
			if i == j {
				dp[i][j] = 1
			}
			if i < j && i+1 < l && j-1 >= 0 {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1] + 2
				} else {
					dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				}
			}
			j++
		}
	}
	return dp[0][l-1]
}
