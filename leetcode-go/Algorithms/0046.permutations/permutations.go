package problem0046

func permute(nums []int) [][]int {
	var result [][]int
	dfs(nums, []int{}, &result)
	return result
}
func dfs(nums, list []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, list)
	}
	for i := range nums {
		dfs(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(list, nums[i]), result)
	}
}
