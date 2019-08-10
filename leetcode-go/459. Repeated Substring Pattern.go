package main

import (
	"fmt"
)

func repeatedSubstringPattern(s string) bool {
	if s == "" {
		return false
	}
	sLen := len(s)
	minLen := 1
	for {
		if sLen%minLen != 0 {
			minLen++
			continue
		}
		if minLen >= sLen {
			return false
		}
		base := s[:minLen]
		// fmt.Println("b:", base)
		skip := false
		for i := 1; i < sLen/minLen; i++ {
			// fmt.Println("t:", s[minLen*i:minLen*(i+1)], minLen*i, minLen*(i+1))
			if base != s[minLen*i:minLen*(i+1)] {
				minLen++
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		return true
	}
}
func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabd"))
}
