package problem0060

import (
	"strconv"
	"strings"
)

var list []string

type cnt struct {
	c int
	k int
}

func getPermutation(n int, k int) string {
	var base []int
	for i := 1; i <= n; i++ {
		base = append(base, i)
	}
	cnt := &cnt{k: k}
	c := make(chan string, 1)
	for _, v := range base {
		dsf(base, strconv.Itoa(v), cnt, c)
	}
	return <-c
}

func dsf(l []int, p string, cnt *cnt, c chan string) {
	if len(p) == len(l) {
		cnt.c++
		if cnt.c == cnt.k {
			c <- p
			return
		}
	}
	for _, v := range l {
		s := strconv.Itoa(v)
		if strings.Index(p, s) == -1 {
			dsf(l, p+s, cnt, c)
		}
	}
}
