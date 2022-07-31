package _022_07_10

import (
	"fmt"
	"strings"
	"testing"
)

func canChange(start string, target string) bool {
	ts, tt := strings.Replace(start, "_", "", -1), strings.Replace(target, "_", "", -1)
	if ts != tt {
		return false
	}
	sl, sr := getRuneIdx(start)
	tl, tr := getRuneIdx(target)
	for i := range sl {
		if tl[i] > sl[i] {
			return false
		}
	}
	for i := range sr {
		if tr[i] < sr[i] {
			return false
		}
	}
	return true
}

func getRuneIdx(s string) (l []int, r []int) {
	for i, v := range s {
		if v == 'L' {
			l = append(l, i)
		}
		if v == 'R' {
			r = append(r, i)
		}
	}
	return l, r
}

func TestCanChange(t *testing.T) {
	fmt.Println(canChange("_L__R__R_", "L______RR"))
	fmt.Println(canChange("R_L_", "__LR"))
	fmt.Println(canChange("_R", "R_"))
}
