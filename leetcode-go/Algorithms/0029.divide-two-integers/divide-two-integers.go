package problem0029

func divide(dividend int, divisor int) int {
	result := int64(dividend) / int64(divisor)
	if result > 1<<31-1 {
		return 1<<31 - 1
	}
	if result < -1<<31 {
		return -1 << 31
	}
	return int(result)
}
