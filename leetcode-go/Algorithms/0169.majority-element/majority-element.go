package problem0169

func majorityElement_1(nums []int) int {
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

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, most := 0, 0
	for _, v := range nums {
		if sum == 0 {
			sum = 1
			most = v
			continue
		}
		if most == v {
			sum++
		} else {
			sum--
		}
	}
	return most
}
