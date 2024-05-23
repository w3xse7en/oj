package _010_minimumCost

import "sort"

func minimumCost(nums []int) int {
	f := nums[0]
	sort.Ints(nums)
	has := false
	for i := 0; i < 3; i++ {
		if nums[i] == f {
			has = true
			break
		}
	}
	if has {
		return nums[0] + nums[1] + nums[2]
	}
	return nums[0] + nums[1] + f
}
