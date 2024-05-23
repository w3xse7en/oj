package _055_canJump

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	mx := -1
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			return true
		}
		if mx <= nums[i] {
			mx = nums[i]
		} else {
			mx--
		}
		if mx == 0 && nums[i] == 0 {
			return false
		}
	}
	return false
}
