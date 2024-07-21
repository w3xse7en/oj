package _151_reverseWords

import (
	"strings"
)

func reverseWords(s string) string {
	fields := strings.Fields(s)
	for i := 0; i < len(fields)/2; i++ {
		fields[i], fields[len(fields)-1-i] = fields[len(fields)-1-i], fields[i]
	}
	return strings.Join(fields, " ")
}
