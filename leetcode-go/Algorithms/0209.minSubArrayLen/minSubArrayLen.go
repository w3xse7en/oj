package _209_minSubArrayLen

import "math"

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	result := math.MaxInt
	for i, j := 0, 0; i < len(nums) && j < len(nums) && i <= j; {
		if sum >= target {
			result = min(result, j-i)
			sum -= nums[i]
			i++
		} else if sum < target {
			j++
			if j < len(nums) {
				sum += nums[j]
			}
		}
		if i > j {
			j = i
			if j < len(nums) {
				sum = nums[i]
			}
		}
	}
	if result == math.MaxInt {
		return 0
	}
	return result + 1
}
