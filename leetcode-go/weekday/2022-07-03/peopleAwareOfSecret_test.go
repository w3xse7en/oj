package _022_07_03

import (
	"fmt"
	"testing"
)

func peopleAwareOfSecret(n int, delay int, forget int) int {
	var mod int = 1e9 + 7
	f := make([]int, n)
	f[0] = 1
	for i := 1; i < n; i++ {
		add := 0
		for j := delay; j < forget; j++ {
			if i-j >= 0 {
				add += f[i-j]
			}
		}
		add %= mod
		f[i] = add
	}
	sum := 0
	for i := 1; i <= forget; i++ {
		sum += f[len(f)-i]
	}
	return sum % mod
}

func TestPeopleAwareOfSecret(t *testing.T) {
	fmt.Println(peopleAwareOfSecret(6, 2, 4))
	fmt.Println(peopleAwareOfSecret(4, 1, 3))
	fmt.Println(peopleAwareOfSecret(684, 18, 496))
}
