package problem0081

func search(nums []int, target int) bool {
	for i, j := 0, len(nums)-1; i <= j; {
		mid := (i + j) / 2
		if nums[mid] == target {
			return true
		}
		if nums[i] == nums[mid] && nums[mid] == nums[j] {
			i++
			j--
			continue
		}
		if nums[mid] >= nums[i] {
			if target >= nums[i] && target <= nums[mid] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		} else if nums[j] >= nums[mid] {
			if target >= nums[mid] && target <= nums[j] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
	}
	return false
}
