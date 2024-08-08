package problem0080

func removeDuplicates(nums []int) int {
	var pre, cnt, idx int
	for i := 0; i < len(nums); i++ {
		if pre == nums[i] {
			cnt++
		} else {
			cnt = 1
			pre = nums[i]
		}
		if cnt <= 2 {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
