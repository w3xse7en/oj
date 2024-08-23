package problem0006

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows == 0 {
		return s
	}
	rows := make([][]rune, numRows)
	i, flag := 0, 0
	for _, r := range s {
		rows[i] = append(rows[i], r)
		if i == numRows-1 {
			flag = -1
		} else if i == 0 {
			flag = 1
		}
		i += flag
	}
	var result strings.Builder
	for _, row := range rows {
		result.WriteString(string(row))
	}
	return result.String()
}
