package problem0003

import "strings"

func lengthOfLongestSubstring(s string) int {
	var n string
	var max int
	for _, v := range s {
		index := strings.Index(n, string(v))
		if index == -1 {
			n += string(v)
		} else {
			n = n[index+1:] + string(v)
		}
		if len(n) >= max {
			max = len(n)
		}
	}
	return max
}
