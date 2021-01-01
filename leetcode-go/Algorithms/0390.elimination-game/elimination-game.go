package problem0390

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	} else {
		return 2 * (n/2 + 1 - lastRemaining(n/2))
	}
}
