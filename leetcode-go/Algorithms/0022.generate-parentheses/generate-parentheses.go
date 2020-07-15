package problem0022

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	r := []string{}
	next := []string{}
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteRune('(')
		sb.WriteRune(')')
	}
	s := sb.String()
	r = append(r, s)
	next = append(next, s)
	for i := 0; i < len(r); i++ {
		f := f(r[i])
		if len(f) == 0 {
			break
		} else {
			r = append(r, f...)
		}
	}
	for _, s := range r {
		fmt.Println(s)
	}

	return r
}

func f(str string) []string {
	r := []string{}
	for i := 0; i < len(str)-1; i++ {
		if str[i] == ')' && str[i+1] == '(' {
			s := []rune(str)
			s[i], s[i+1] = s[i+1], s[i]
			r = append(r, string(s))
		}
	}
	return r
}

func is(s string) bool {
	t := []rune(s)
	for i, j := 0, len(s); i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
	return string(t) == s
}
