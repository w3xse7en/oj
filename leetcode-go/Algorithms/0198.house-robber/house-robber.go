package problem0198

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 0; i < len(nums); i++ {
		t2, t3 := 0, 0
		if i-2 >= 0 {
			t2 = nums[i-2]
		}
		if i-3 >= 0 {
			t3 = nums[i-3]
		}
		t := t2
		if t3 > t2 {
			t = t3
		}
		nums[i] += t
	}
	result := nums[len(nums)-1]
	if nums[len(nums)-2] > result {
		result = nums[len(nums)-2]
	}
	return result
}
