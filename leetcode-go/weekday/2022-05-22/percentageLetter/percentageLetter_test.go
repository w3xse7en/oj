package percentageLetter

import (
	"fmt"
	"testing"
)

func percentageLetter(s string, letter byte) int {
	if len(s) == 0 {
		return 0
	}
	var cnt int
	for _, v := range s {
		if v == rune(letter) {
			cnt++
		}
	}
	return cnt * 100 / len(s)
}

func TestA(t *testing.T) {
	fmt.Println(percentageLetter("foobar", 'o'))
	fmt.Println(percentageLetter("", 'o'))
	fmt.Println(percentageLetter("", '1'))
	fmt.Println(percentageLetter("1111", 0))
	fmt.Println(percentageLetter("", 0))
}
