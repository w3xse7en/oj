package problem0167

import "slices"

func twoSum(numbers []int, target int) []int {
	for i := range numbers {
		index := slices.Index(numbers[i+1:], target-numbers[i])
		if index != -1 {
			return []int{i + 1, index + i + 1 + 1}
		}
	}
	return nil
}
