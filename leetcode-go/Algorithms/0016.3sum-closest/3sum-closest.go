package problem0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	var (
		min     = float64(0xffffffff)
		closest int
		result  int
	)
	for i+2 <= j {
		for ; i+2 <= j; i++ {
			t := target - nums[i] - nums[j]
			one := oneClosest(nums[i+1:j], t)
			tm := math.Abs(float64(t - one))
			closest = nums[i] + nums[j] + one
			if tm < min {
				min = tm
				result = closest
			}
			if closest == target {
				return closest
			}
		}
		i = 0
		j--
	}
	return result
}

func oneClosest(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) / 2
		m := nums[mid]
		if target == m {
			return target
		} else if m < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	closest := nums[i]
	min := math.Abs(float64(nums[i] - target))
	if i+1 > 0 && i+1 < len(nums) {
		i1 := math.Abs(float64(nums[i+1] - target))
		if i1 < min {
			min = i1
			closest = nums[i+1]
		}
	}
	if i-1 > 0 && i-1 < len(nums) {
		i1 := math.Abs(float64(nums[i-1] - target))
		if i1 < min {
			min = i1
			closest = nums[i-1]
		}
	}
	return closest
}
