package problem0053

// dp[i] = max(dp[i],dp[i-1]+nums[i])
func maxSubArray(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i-1]+nums[i])
		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
