package problem0217

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for k, num := range nums {
		if k > 0 && num == nums[k-1] {
			return true
		}
	}
	return false
}
