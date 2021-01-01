package problem0072

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1) + 1
	n := len(word2) + 1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			for j := 1; j < n; j++ {
				dp[i][j] = j
			}
		}
		dp[i][0] = i
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf(" %v", dp[i][j])
		}
		fmt.Println()
	}
	fmt.Println("----")

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
