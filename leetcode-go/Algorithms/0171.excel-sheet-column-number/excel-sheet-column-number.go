package problem0171

func titleToNumber(s string) int {
	result, step := 0, 1
	for i := len(s) - 1; i >= 0; i-- {
		result += int(s[i]-'A'+1) * step
		step *= 26
	}
	return result
}
