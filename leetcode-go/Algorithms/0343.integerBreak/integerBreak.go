package _343_integerBreak

import (
	"math"
)

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	mx := 0
	sqrt := int(math.Sqrt(float64(n))) + 1
	for i := 2; i <= sqrt; i++ {
		m := n % i
		t := 0
		switch m {
		case 0:
			t = pow(i, n/i)
		case 1:
			if (n/i)-1 > 0 {
				t = pow(i, (n/i)-1) * (i + 1)
			} else {
				t = i
			}
		default:
			t = pow(i, n/i) * m
		}
		if t > mx {
			mx = t
		}
	}
	return mx
}

func pow(a, b int) int {
	c := a
	for i := 0; i < b-1; i++ {
		c *= a
	}
	return c
}
