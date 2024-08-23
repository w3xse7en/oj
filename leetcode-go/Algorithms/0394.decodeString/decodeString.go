package _394_decodeString

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	var result string
	var tmpNumb string
	var ints []int
	for _, r := range s {
		if unicode.IsNumber(r) {
			tmpNumb += string(r)
			continue
		}
		if r == '[' {
			n, _ := strconv.Atoi(tmpNumb)
			ints = append(ints, n)
			tmpNumb = ""
			result += "0"
			continue
		}
		if r == ']' {
			idx := 0
			for i := len(result) - 1; i >= 0; i-- {
				if result[i] == '0' {
					idx = i
					break
				}
			}
			newS := build(result[idx+1:], ints[len(ints)-1])
			ints = ints[:len(ints)-1]
			result = result[:idx] + newS
			continue
		}
		result = result + string(r)
	}
	return result
}

func build(s string, n int) string {
	var str strings.Builder
	for i := 0; i < n; i++ {
		str.WriteString(s)
	}
	return str.String()
}
