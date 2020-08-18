package problem0034

func searchRange(nums []int, target int) []int {
	for i, j := 0, len(nums)-1; ; {
		if i > j {
			return []int{-1, -1}
		}
		mid := (i + j) / 2
		if target == nums[mid] {
			return []int{fl(nums, target, i, mid), fr(nums, target, mid, j)}
		} else if target > nums[mid] {
			i = mid + 1
		} else if target < nums[mid] {
			j = mid - 1
		}
	}
}
func fl(nums []int, target, i, j int) int {
	for {
		if i == j {
			return i
		}
		mid := (i + j) / 2
		if target == nums[mid] {
			j = mid
		} else {
			i++
		}
	}
}
func fr(nums []int, target, i, j int) int {
	for {
		if i == j {
			return i
		}
		mid := (i + j + 1) / 2
		if target == nums[mid] {
			i = mid
		} else {
			j--
		}
	}
}
