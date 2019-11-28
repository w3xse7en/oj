package problem0066

func plusOne(digits []int) []int {
	ld := len(digits)
	plus := false
	for i := ld - 1; i >= 0; i-- {
		if i == ld-1 {
			plus = true
		}
		if plus {
			digits[i]++
			plus = false
		}
		if digits[i] == 10 {
			digits[i] = 0
			plus = true
		}
	}
	if plus {
		digits = append([]int{1}, digits...)
	}
	return digits
}
