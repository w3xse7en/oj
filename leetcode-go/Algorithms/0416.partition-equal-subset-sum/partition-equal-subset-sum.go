package problem0416

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	sum /= 2

	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		if i == 0 {
			dp[0][0] = true
		}
		if i >= 1 {
			for j := 1; j < len(dp[i]); j++ {
				if j-nums[i-1] < 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
				}
			}
		}
	}
	return dp[len(nums)][sum]
}
