package _035_maxUncrossedLines

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	mx := 0
	for i, v1 := range nums1 {
		dp[i] = make([]int, len(nums2))
		for j, v2 := range nums2 {
			a, b, c := 0, 0, 0
			if i-1 >= 0 && j-1 >= 0 {
				a = dp[i-1][j-1]
			}
			if i-1 >= 0 {
				b = dp[i-1][j]
			}
			if j-1 >= 0 {
				c = dp[i][j-1]
			}
			if v1 == v2 {
				dp[i][j] = a + 1
			} else {
				dp[i][j] = max(b, c)
			}
			mx = max(mx, dp[i][j])
		}
	}
	return mx
}
