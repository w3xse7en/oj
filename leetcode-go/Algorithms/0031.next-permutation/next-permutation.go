package problem0031

import "sort"

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		t, m := i, 0xffffffff
		for j := i + 1; j < len(nums); j++ {
			c := nums[j] - nums[i]
			if c > 0 && c < m {
				m = c
				t = j
			}
		}
		if t != i {
			nums[i], nums[t] = nums[t], nums[i]
			sort.Ints(nums[i+1:])
			return
		}
	}
	sort.Ints(nums)
	return
}
