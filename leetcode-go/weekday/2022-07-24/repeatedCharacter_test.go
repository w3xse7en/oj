package _022_07_24

import "testing"

func repeatedCharacter(s string) byte {
	mp := map[byte]bool{}
	for _, v := range s {
		_, ok := mp[byte(v)]
		if !ok {
			mp[byte(v)] = true
		} else {
			return byte(v)
		}
	}
	return 0
}

func TestRepeatedCharacter(t *testing.T) {

}
