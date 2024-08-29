package _224_calculate

import (
	"strconv"
	"strings"
	"unicode"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	list := []string{}
	numb := []rune{}
	var pre rune
	for _, r := range s {
		if unicode.IsNumber(r) {
			numb = append(numb, r)
			continue
		}
		if len(numb) != 0 {
			n := string(numb)
			if pre == '-' {
				n = "-" + n
			}
			list = append(list, n)
			numb = []rune{}
		}
		if r == '(' {
			list = append(list, string(r))
			if pre == '-' {
				list = append(list, "-")
			}
		}
		if r == ')' {
			leftIdx := 0
			sum := 0
			for i := len(list) - 1; i >= 0; i-- {
				v := list[i]
				if v == "(" {
					leftIdx = i
					break
				}
				if v == "-" {
					sum *= -1
					continue
				}
				n, _ := strconv.Atoi(v)
				sum += n
			}
			list = append(list[:leftIdx], strconv.Itoa(sum))
		}
		if !unicode.IsNumber(r) {
			pre = r
		}
	}
	if len(numb) != 0 {
		n := string(numb)
		if pre == '-' {
			n = "-" + n
		}
		list = append(list, n)
		numb = []rune{}
	}
	sum := 0
	if len(list) != 0 {
		for i := len(list) - 1; i >= 0; i-- {
			v := list[i]
			if v == "-" {
				sum *= -1
				continue
			}
			n, _ := strconv.Atoi(v)
			sum += n
		}
	}
	return sum
}
