package _022_07_31

import "testing"

func TestMaximumGroups(t *testing.T) {

}

func maximumGroups(grades []int) int {
	step := 2
	ans := 1
	length := 3
	for len(grades) >= length {
		ans += 1
		step += 1
		length += step
	}
	return ans
}
