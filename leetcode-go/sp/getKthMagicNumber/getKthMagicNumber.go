package getKthMagicNumber

func getKthMagicNumber(k int) int {
	dp := make([]int, k)
	dp[0] = 1
	a, b, c := 0, 0, 0
	for i := 1; i < k; i++ {
		xa, xb, xc := dp[a]*3, dp[b]*5, dp[c]*7
		dp[i] = min(xa, min(xb, xc))
		if dp[i] == xa {
			a++
		}
		if dp[i] == xb {
			b++
		}
		if dp[i] == xc {
			c++
		}
	}
	return dp[k-1]
}
