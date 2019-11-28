package problem0070

// Mind
// 1
// 1
// sum 1

// 2
// 1 1
// 2
// sum 2

// 3
// 1 1 1
// 1 2
// 2 1
// sum = 1+2 = 3

// 4
// 1 1 1 1
// 1 1 2
// 1 2 1
// 2 1 1
// 2 2
// sum = 3+2 = 5

// 5
// 1 1 1 1 1
// 1 1 1 2
// 1 1 2 1
// 1 2 1 1
// 2 1 1 1
// 1 2 2
// 2 1 2
// 2 2 1
// sum = 5+3 = 8
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b := 1, 2
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
