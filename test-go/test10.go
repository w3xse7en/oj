package main

import (
	"fmt"
	"strings"
)

func STest(sent []string, query []string) {
	for _, v := range query {
		qy := strings.Split(v, " ")
		first := true
		fail := true
		for k, sv := range sent {
			cnt := 0
			for _, qv := range qy {
				if !strings.Contains(sv, qv) {
					cnt++
				}
			}
			if cnt == len(qy) {
				fail = false
				if first {
					fmt.Print(k)
					first = false
				} else {
					fmt.Print(" ", k)
				}
			}
		}
		if fail {
			fmt.Println(-1)
		} else {
			fmt.Println()
		}
	}
}

func main() {
	STest([]string{"a b c", "b c", "c d e"}, []string{"b", "c d", "e", "f"})
}
