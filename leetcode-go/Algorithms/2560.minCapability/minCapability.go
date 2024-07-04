package _560_minCapability

import "math"

var mi = math.MaxInt

func minCapability(nums []int, k int) int {
	mi = math.MaxInt
	loop(nums, 1, k, 0)
	return mi
}

func loop(nums []int, dep, k, mx int) {
	for i := 0; i < len(nums)-2*(k-dep); i++ {
		tmx := mx
		if nums[i] > tmx {
			tmx = nums[i]
		}
		if dep < k && i+2 < len(nums) {
			loop(nums[i+2:], dep+1, k, tmx)
		}
		if dep == k {
			if tmx < mi {
				mi = tmx
			}
		}
	}
}
