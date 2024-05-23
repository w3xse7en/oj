package problem0015

import (
	"sort"
)

func threeSum1(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	// fmt.Println(nums)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j] > 0 {
				break
			}
			for k := j + 1; k < len(nums); k++ {
				if k != j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				} else if nums[i]+nums[j]+nums[k] > 0 {
					break
				}
			}
		}
	}
	return result
}
