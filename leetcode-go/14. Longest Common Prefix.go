package main

import (
	"fmt"
)

func longestCommonPrefixOwn(strs []string) string {
	short := ""
	for k, long := range strs {
		if k == 0 {
			short = long
		} else {
			if len(short) > len(long) {
				short, long = long, short
			}
			for i := 0; i < len(short); i++ {
				if short[i:i+1] != long[i:i+1] {
					short = short[0:i]
				}
			}
		}
	}
	return short
}

func main() {
	fmt.Println("r:", longestCommonPrefixOwn([]string{"flower", "flow", "flight"}))
	fmt.Println("r:", longestCommonPrefixOwn([]string{"dog", "racecar", "car"}))
	fmt.Println("r:", longestCommonPrefixOwn([]string{"aaaaaa", "aaaaaaaaaa", "caaaaar"}))
}

// O(N)
func longestCommonPrefixBetter(strs []string) string {
	var min, max string
	for _, str := range strs {
		if str == "" {
			return ""
		}
		if min == "" || min > str {
			min = str
		}
		if max == "" || max < str {
			max = str
		}
	}
	var result []byte
	for i := 0; i < len(min); i++ {
		if len(max) <= i || min[i] != max[i] {
			return string(result)
		}
		result = append(result, min[i])
	}

	return string(result)
}
