package problem0172

func trailingZeroes(n int) int {
	r := []int{1}
	for i := 2; i <= n; i++ {
		r = multiple(r, iToArray(i))
	}
	cnt := 0
	for _, v := range r {
		if v == 0 {
			cnt++
		} else {
			break
		}
	}
	return cnt
}

func iToArray(n int) []int {
	r := []int{}
	for ; n > 0; n /= 10 {
		r = append(r, n%10)
	}
	return r
}

func multiple(a []int, b []int) []int {
	s := make([][]int, len(a))
	for ka, va := range a {
		for i := 0; i < ka; i++ {
			s[ka] = append(s[ka], 0)
		}
		for _, vb := range b {
			s[ka] = append(s[ka], va*vb)
		}
	}

	maxLen := len(s[len(s)-1])
	r := []int{}
	step := 0
	for i := 0; i < maxLen; i++ {
		var z int
		for _, v := range s {
			if i > len(v)-1 {
				z += 0
			} else {
				z += v[i]
			}
		}
		z += step
		if z >= 10 {
			step = z / 10
			z %= 10
		} else {
			step = 0
		}
		r = append(r, z)
	}
	if step != 0 {
		r = append(r, step)
	}
	return r
}
