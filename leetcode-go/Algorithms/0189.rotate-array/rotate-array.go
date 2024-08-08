package problem0189

func rotate1(nums []int, k int) {
	l := len(nums)
	for k, v := range append(nums[l-k%l:], nums[:l-k%l]...) {
		nums[k] = v
	}
	return
}
