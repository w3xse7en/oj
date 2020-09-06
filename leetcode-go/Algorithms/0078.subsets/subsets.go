package problem0078

func subsets(nums []int) [][]int {
	var result [][]int
	for k := range nums {
		dsf(k, nums, []int{}, &result)
	}
	result = append(result, []int{})
	return result
}

func dsf(k int, nums, tmp []int, result *[][]int) {
	if k < len(nums) {
		tmp = append(tmp, nums[k])
		t := make([]int, len(tmp))
		copy(t, tmp)
		*result = append(*result, t)
	}
	for i := k + 1; i < len(nums); i++ {
		dsf(i, nums, tmp, result)
	}
}
