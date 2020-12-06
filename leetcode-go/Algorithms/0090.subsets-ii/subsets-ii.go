package problem0090

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var noDup map[string]bool

func subsetsWithDup(nums []int) [][]int {
	var list [][]int
	if len(nums) == 0 {
		return list
	}
	sort.Ints(nums)
	noDup = make(map[string]bool)
	list = append(list, []int{})
	loop(nums, []int{}, &list)
	fmt.Printf("%+v", list)
	return list
}

func loop(nums, l []int, list *[][]int) {
	for i := 0; i < len(nums); i++ {
		*list = noDupAppend(*list, append(l, nums[i]))
		if i+1 < len(nums) {
			loop(nums[i+1:], append(l, nums[i]), list)
		}
	}
}
func noDupAppend(list [][]int, result []int) [][]int {
	var sb strings.Builder
	l := make([]int, len(result))
	for k, v := range result {
		sb.WriteString(strconv.Itoa(v))
		l[k] = v
	}
	if noDup[sb.String()] {
		return list
	}
	noDup[sb.String()] = true
	return append(list, l)
}
