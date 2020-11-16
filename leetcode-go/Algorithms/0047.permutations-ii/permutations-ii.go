package problem0047

func permuteUnique(nums []int) [][]int {
	var unique [][]int
	loop(0, nums, &unique)
	return unique
}
func loop(idx int, nums []int, unique *[][]int) {
	if idx == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*unique = append(*unique, tmp)
	}
	visited := make(map[int]bool, 0)
	for i := idx; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		nums[i], nums[idx] = nums[idx], nums[i]
		loop(idx+1, nums, unique)
		nums[i], nums[idx] = nums[idx], nums[i]
		visited[nums[i]] = true
	}
}
