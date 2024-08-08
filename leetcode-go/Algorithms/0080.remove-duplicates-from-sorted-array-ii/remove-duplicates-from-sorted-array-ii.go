package problem0080

func removeDuplicates1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	idx := 2
	for i := 2; i < len(nums); i++ {
		if nums[idx-2] != nums[i] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
