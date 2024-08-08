package _264_nthUglyNumber

import (
	"fmt"
	"sort"
)

func nthUglyNumber(n int) int {
	list := []int{1}
	exist := map[int]bool{}
	for i := 0; i < n; i++ {
		a := list[i] * 2
		b := list[i] * 3
		c := list[i] * 5
		if !exist[a] {
			exist[a] = true
			list = append(list, a)
		}
		if !exist[b] {
			exist[b] = true
			list = append(list, b)
		}
		if !exist[c] {
			exist[c] = true
			list = append(list, c)
		}
		fmt.Println(list)
		sort.Ints(list)
	}
	fmt.Println(list[n-1])
	return list[n-1]
}
