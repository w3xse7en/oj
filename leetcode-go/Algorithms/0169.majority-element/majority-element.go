package problem0169

func majorityElement(nums []int) int {
	mp := map[int]int{}
	t := len(nums) / 2
	for _, v := range nums {
		mp[v]++
		if mp[v] > t {
			return v
		}
	}
	return 0
}
