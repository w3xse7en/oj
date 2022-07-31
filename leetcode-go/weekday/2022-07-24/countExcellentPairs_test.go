package _022_07_24

import (
	"fmt"
	"math/bits"
	"strings"
	"testing"
)

type p struct {
	a int
	b int
}

func countExcellentPairs(nums []int, k int) int64 {
	exist := map[int]bool{}
	count := map[int]int{}
	for _, num := range nums {
		if exist[num] {
			continue
		}
		exist[num] = true
		count[bits.OnesCount(uint(num))]++
	}
	var total int64
	for a, aa := range count {
		for b, bb := range count {
			if a+b >= k {
				total += int64(aa * bb)
			}
		}
	}
	return total
}

func is(a int, b int, k int, exist map[p]bool) (int64, map[p]bool) {
	tp := p{
		a: a,
		b: b,
	}
	if exist[tp] {
		return 0, exist
	} else {
		exist[tp] = true
	}
	and := fmt.Sprintf("%b", a&b)
	or := fmt.Sprintf("%b", a|b)
	one := 0
	one += strings.Count(and, "1")
	one += strings.Count(or, "1")
	if one >= k {
		//fmt.Printf("%v %v %b,%b\n", a, b, a&b, a|b)
		return 1, exist
	}
	return 0, exist
}

func TestCountExcellentPairs(t *testing.T) {
	fmt.Println(countExcellentPairs([]int{1, 2, 3, 1}, 3))

}
