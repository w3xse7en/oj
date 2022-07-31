package _022_07_03

import (
	"fmt"
	"testing"
)

func decodeMessage(key string, message string) string {
	replace := map[rune]rune{}
	st := 'a'
	for _, v := range key {
		if v >= 'a' && v <= 'z' {
			if _, ok := replace[v]; !ok {
				replace[v] = st
				st++
			}
		}
	}
	var newMsg []rune
	for _, v := range message {
		if _, ok := replace[v]; ok {
			newMsg = append(newMsg, replace[v])
		} else {
			newMsg = append(newMsg, v)
		}
	}
	return string(newMsg)
}

func TestDecodeMessage(t *testing.T) {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
	fmt.Println(decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))
}
