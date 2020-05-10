package problem0191

import (
	"strconv"
)

func hammingWeight(num uint32) int {
	var cnt int
	for _, v := range strconv.FormatUint(uint64(num), 2) {
		if v == '1' {
			cnt++
		}
	}
	return cnt
}
