package problem0015

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			t := -(nums[i] + nums[j])
			if contains(nums[j+1:], t) {
				result = append(result, []int{nums[i], nums[j], t})
			}
		}
	}
	return result
}

func contains(nums []int, t int) bool {
	i, j := 0, len(nums)-1
	for i <= j {
		h := (i + j) / 2
		if nums[h] > t {
			j = h - 1
		} else if nums[h] < t {
			i = h + 1
		} else {
			return true
		}
	}
	return false
}
