package problem0204

func countPrimes1(n int) int {
	if n < 2 {
		return 0
	}
	list := make([]bool, n)
	list[0] = true
	list[1] = true
	for i := 2; i*i < n; i++ {
		for j := 2; j*i < n; j++ {
			list[i*j] = true
		}
	}
	var sum int
	for _, ok := range list {
		if !ok {
			sum++
		}
	}
	return sum
}
