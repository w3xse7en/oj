package problem0686

import (
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	var base strings.Builder
	base.WriteString(A)
	cnt := 1
	for {
		if base.Len() < len(B) {
			base.WriteString(A)
			cnt++
			continue
		}
		if strings.Contains(base.String(), B) {
			return cnt
		}
		base.WriteString(A)
		cnt++
		if base.Len() > 10000 {
			return -1
		}
	}
}
