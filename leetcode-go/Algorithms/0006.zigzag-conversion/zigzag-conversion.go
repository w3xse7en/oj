package problem0006

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows <= 1 {
		return s
	}
	c := make([][]rune, numRows)
	line, step := 0, 0
	for _, v := range s {
		c[line] = append(c[line], v)
		if line == numRows-1 {
			step = -1
		}
		if line == 0 {
			step = 1
		}
		line += step
	}
	var result []rune
	for _, runes := range c {
		result = append(result, runes...)
	}
	return string(result)
}
