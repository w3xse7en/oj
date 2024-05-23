package _045_jump

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	mx := nums[0]
	step := 0
	j := nums[0]
	for i := 0; i < len(nums); i++ {
		mx = max(nums[i], mx-1)
		if j <= 0 || i == len(nums)-1 {
			step++
			j = mx
		}
		j--
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
