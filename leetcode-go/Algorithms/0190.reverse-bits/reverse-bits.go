package problem0190

import (
	"fmt"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	formatUint := fmt.Sprintf("%032v", strconv.FormatUint(uint64(num), 2))
	parseUint, err := strconv.ParseUint(Reverse(formatUint), 2, 64)
	if err != nil {
		return 0
	}
	return uint32(parseUint)
}
func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
