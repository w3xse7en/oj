package main

import (
	"fmt"
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

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))     // 3
	fmt.Println(repeatedStringMatch("abcd", "bcdab"))        // 2
	fmt.Println(repeatedStringMatch("abc", "cabcabca"))      // 4
	fmt.Println(repeatedStringMatch("abcabcabcabc", "abac")) // -1
}
