package problem0204

func countPrimes(n int) int {
	cnt := 0
	list := make([]bool, n)
	for i := 2; i < n; i++ {
		if list[i] == false {
			cnt++
			for j := i; j < n; j += i {
				list[j] = true
			}
		}
	}
	return cnt
}
