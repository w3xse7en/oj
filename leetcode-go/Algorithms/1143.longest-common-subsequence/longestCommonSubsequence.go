package problem1143

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := range text1 {
		dp[i] = make([]int, len(text2))
		for j := i; j < len(text2); j++ {
			a, b, c := 0, 0, 0
			if i-1 >= 0 {
				a = dp[i-1][j]
			}
			if j-1 >= 0 {
				b = dp[i][j-1]
			}
			if i-1 >= 0 && j-1 >= 0 {
				c = dp[i-1][j-1]
			}
			dp[i][j] = max(b, c)
			if text1[i] == text2[j] {
				dp[i][j] = a + 1
			}
		}
	}
	return dp[len(text1)-1][len(text2)-1]
}
