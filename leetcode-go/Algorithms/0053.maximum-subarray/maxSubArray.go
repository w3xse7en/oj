package problem0053

func maxSubArray(nums []int) int {
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i]+nums[i-1])
		mx = max(nums[i], mx)
	}
	return mx
}
