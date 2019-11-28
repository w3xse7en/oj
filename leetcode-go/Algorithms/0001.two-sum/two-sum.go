package problem0001

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for indices, num := range nums {
		anotherIndices, ok := hash[target-num]
		if ok {
			return []int{indices, anotherIndices}
		}
		hash[num] = indices
	}
	return nil
}
