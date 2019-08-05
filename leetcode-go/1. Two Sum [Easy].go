package main

import (
	"fmt"
)

// Given nums = [2, 7, 11, 15], target = 9,
//
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

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

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 5}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
