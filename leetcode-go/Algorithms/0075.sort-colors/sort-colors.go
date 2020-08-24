package problem0075

func sortColors(nums []int) {
	zero, two := 0, len(nums)-1
	for i := zero; i <= two; i++ {
		if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		} else if nums[i] == 2 {
			nums[i], nums[two] = nums[two], nums[i]
			two--
			i--
		}
	}
	return
}
