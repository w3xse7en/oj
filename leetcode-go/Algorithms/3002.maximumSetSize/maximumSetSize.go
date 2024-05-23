package _002_maximumSetSize

func maximumSetSize(nums1 []int, nums2 []int) int {
	mp1, mp2 := map[int]bool{}, map[int]bool{}

	for _, v := range nums1 {
		mp1[v] = true
	}
	for _, v := range nums2 {
		if len(mp1) <= len(nums1)/2 {
			if !mp1[v] {
				mp2[v] = true
			}
		} else {
			if !mp2[v] {
				mp2[v] = true
				if mp1[v] {
					delete(mp1, v)
				}
			}
		}
	}
	l1 := len(mp1)
	if l1 > len(nums1)/2 {
		l1 = len(nums1) / 2
	}
	l2 := len(mp2)
	if l2 > len(nums2)/2 {
		l2 = len(nums2) / 2
	}
	return l1 + l2
}
