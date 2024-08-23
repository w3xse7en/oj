package _007_findMaximumNumber

import (
	"fmt"
	"math"
)

func findMaximumNumber(k int64, x int) int64 {
	var sum int64
	var idxList []int
	for i := 1; i < 65; i++ {
		idx := x*i - 1
		if idx < 65 {
			idxList = append(idxList, idx)
		}
	}
	for i := int64(1); i < math.MaxInt; i++ {
		binary := fmt.Sprintf("%0b", i)
		for _, idx := range idxList {
			if idx < len(binary) && binary[len(binary)-1-idx] == '1' {
				sum += 1
			}
		}
		if sum > k {
			return i - 1
		}
	}
	return 0
}
