package _95_longestSubstring

import "strings"

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	char := map[rune]int{}
	for _, v := range s {
		char[v]++
	}
	for r, cnt := range char {
		l := 0
		if cnt < k {
			for _, s := range strings.Split(s, string(r)) {
				l = m(l, longestSubstring(s, k))
			}
		}
		return l
	}
	return len(s)
}
func m(a, b int) int {
	if a > b {
		return a
	}
	return b
}
