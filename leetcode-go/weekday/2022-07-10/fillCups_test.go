package _022_07_10

import (
	"fmt"
	"sort"
	"testing"
)

type am struct {
	idx   int
	value int
}

func fillCups(amount []int) int {
	var second int
	var list = make([]am, len(amount))
	for {
		for k := range amount {
			list[k] = am{
				idx:   k,
				value: amount[k],
			}
		}
		sort.Slice(list, func(i, j int) bool {
			return list[i].value > list[j].value
		})
		if list[0].value > 0 {
			amount[list[0].idx]--
		}
		if list[1].value > 0 {
			amount[list[1].idx]--
		}
		if list[0].value == 0 {
			goto end
		}
		second++
	}
end:
	return second
}

func TestFillCups(t *testing.T) {
	fmt.Println(fillCups([]int{1, 4, 2}))
	fmt.Println(fillCups([]int{5, 4, 4}))
	fmt.Println(fillCups([]int{5, 0, 0}))
}
