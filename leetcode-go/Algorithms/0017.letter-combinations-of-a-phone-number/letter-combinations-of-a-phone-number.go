package problem0017

import (
	"sort"
	"strings"
)

var mp = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	var result []string
	for _, v := range digits {
		var tmp []string
		for _, s := range mp[v] {
			flag := true
			for _, r := range result {
				flag = false
				var sb strings.Builder
				sb.WriteString(r)
				sb.WriteRune(s)
				tmp = append(tmp, sb.String())
			}
			if flag {
				tmp = append(tmp, string(s))
			}
		}
		result = tmp
	}
	sort.Sort(sort.Reverse(sort.StringSlice(result)))
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}
