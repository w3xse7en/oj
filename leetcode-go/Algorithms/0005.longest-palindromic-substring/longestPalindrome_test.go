package problem0005

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 8; i++ {
		m[i] = i - 1
	}
	for i := 0; i < 10; i++ {
		fmt.Println("round", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}
