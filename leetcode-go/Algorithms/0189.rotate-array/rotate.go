package problem0189

import "slices"

func rotate2(nums []int, k int) {
	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}
	for i, v := range append(nums[l-k:], nums[:l-k]...) {
		nums[i] = v
	}
	return
}
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
	return
}
