package _022_07_31

import (
	"fmt"
	"sort"
	"testing"
)

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	tmp := 0
	cnt := 0
	for k := range nums {
		nums[k] -= tmp
		if nums[k] == 0 {
			continue
		}
		tmp += nums[k]
		cnt++
	}
	return cnt
}

func TestMinimumOperations(t *testing.T) {
	//fmt.Println(minimumOperations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(minimumOperations([]int{9, 9, 9, 9}))
	fmt.Println(minimumOperations([]int{9, 8, 9, 9}))
	fmt.Println(minimumOperations([]int{9, 8, 1, 9}))
	fmt.Println(minimumOperations([]int{9, 8, 1, 19}))
	fmt.Println(minimumOperations([]int{0}))
}
