package _153_findMin

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		vl, vr, vm := nums[l], nums[r], nums[mid]
		//fmt.Println(l, vl, "|", r, vr, "|", mid, vm)
		if vm == vl || vm == vr {
			break
		}
		if vl > vm {
			r = mid
			continue
		}
		if vm > vr {
			l = mid
			continue
		}
		r = mid
	}
	return min(nums[l], nums[r])
}
