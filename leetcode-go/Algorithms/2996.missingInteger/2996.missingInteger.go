package _996_missingInteger

import (
	"slices"
	"sort"
)

func missingInteger(nums []int) int {
	var mx = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			mx += nums[i]
		} else {
			break
		}
	}
	sort.Ints(nums)
	index := slices.Index(nums, mx)
	if index == -1 {
		return mx
	}
	for i := index; i < len(nums); i++ {
		if mx == nums[i] {
			mx++
		} else {
			break
		}
	}
	return mx
}
