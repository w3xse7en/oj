package problem0033

func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for {
		if i > j {
			return -1
		}
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
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
}
