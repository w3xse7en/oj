package problem0018

import (
	"errors"
	"sort"
)

// assumes that nums is sorted.
func sumOfTwoWithTarget(nums []int, target int) [][]int {
	numberPartOfASolution := make(map[int]bool)
	solutions := make([][]int, 0, len(nums)/2+1)

	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			solutions = append(solutions, []int{nums[i], nums[j]})
			numberPartOfASolution[nums[i]] = true
			numberPartOfASolution[nums[j]] = true
		}

		for ; i < j && numberPartOfASolution[nums[i]]; i++ {
		}
		for ; i < j && numberPartOfASolution[nums[j]]; j-- {
		}

		if sum < target {
			i++
		} else if sum > target {
			j--
		}
	}

	return solutions
}

func kSum(nums []int, target int, k int) [][]int {
	//fmt.Printf("k = %d ; nums = %#v ; target = %d\n", k, nums, target)
	if k == 2 {
		return sumOfTwoWithTarget(nums, target)
	} else {
		sets := make([][]int, 0)
		searchedSolutionStartingWithN := make(map[int]bool)
		for i := 0; i < len(nums); i++ {
			if searchedSolutionStartingWithN[nums[i]] {
				continue
			}
			searchedSolutionStartingWithN[nums[i]] = true
			lowerSetSolutions := kSum(nums[i+1:], target-nums[i], k-1)
			for _, lowerSet := range lowerSetSolutions {
				sets = append(sets, append(lowerSet, nums[i]))
			}
		}
		return sets
	}
}

func kSumWrapper(nums []int, target int, k int) ([][]int, error) {
	sort.Ints(nums)
	if k < 2 {
		return [][]int{}, errors.New("K has to be atleast 2")
	}
	return kSum(nums, target, k), nil
}

func fourSum(nums []int, target int) [][]int {
	ans, _ := kSumWrapper(nums, target, 4)
	return ans
}
