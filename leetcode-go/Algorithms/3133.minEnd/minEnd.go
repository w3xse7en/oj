package _133_minEnd

import (
	"fmt"
	"strconv"
)

func minEnd(n int, x int) int64 {
	bx, bn := []rune(fmt.Sprintf("%064b", x)), []rune(fmt.Sprintf("%064b", n-1))
	for i, j := len(bn)-1, len(bx)-1; i >= 0 && j >= 0; {
		if bx[j] == '0' {
			bx[j] = bn[i]
			j--
			i--
			continue
		}
		j--
	}
	br, _ := strconv.ParseInt(string(bx), 2, 64)
	return br
}
