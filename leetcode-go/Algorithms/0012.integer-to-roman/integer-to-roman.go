package problem0012

import (
	"strings"
)

var r = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
	4:    "IV",
	9:    "IX",
	40:   "XL",
	90:   "XC",
	400:  "CD",
	900:  "CM",
}

func intToRoman(num int) string {
	rs := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var sb strings.Builder
	for _, v := range rs {
		tmp := num / v
		for i := 0; i < tmp; i++ {
			sb.WriteString(r[v])
		}
		num -= tmp * v
	}
	return sb.String()
}
