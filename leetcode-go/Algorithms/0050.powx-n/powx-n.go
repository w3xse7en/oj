package problem0050

import (
	"strings"
)

func translateNum(num int) int {
	mp := map[string]string{
		"ba": "k", // 10
		"bb": "l",
		"bc": "m",
		"bd": "n",
		"be": "o",
		"bf": "p",
		"bg": "q",
		"bh": "r",
		"bi": "s",
		"bj": "t",
		"ca": "u", // 20
		"cb": "v",
		"cc": "w",
		"cd": "x",
		"ce": "y",
		"cf": "z", //25
	}
	var sb strings.Builder
	for num != 0 {
		n := num % 10
		num /= 10
		sb.WriteByte(byte(n + 'a'))
	}

	st := []rune(sb.String())
	for i, j := 0, len(st)-1; i < j; i, j = i+1, j-1 {
		st[i], st[j] = st[j], st[i]
	}
	sl := []string{string(st)}
	ex := map[string]bool{string(st): true}
	for i := 0; i < len(sl); i++ {
		s := sl[i]
		for j := 0; j < len(s)-1; j++ {
			var sb strings.Builder
			sb.WriteByte(s[j])
			sb.WriteByte(s[j+1])
			is := sb.String()
			s2, ok := mp[is]
			if ok {
				replace := s[:j] + s2 + s[j+2:]
				if !ex[replace] {
					ex[replace] = true
					sl = append(sl, replace)
				}
			}
		}
	}
	return len(sl)
}
