package _043_longestCommonPrefix

import "strconv"

type Node struct {
	mp map[rune]*Node
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	root := &Node{mp: map[rune]*Node{}}
	for _, i := range arr1 {
		node := root
		for _, r := range strconv.Itoa(i) {
			if node.mp[r] == nil {
				node.mp[r] = &Node{map[rune]*Node{}}
			}
			node = node.mp[r]
		}
	}
	var mx int
	for _, i := range arr2 {
		var cnt int
		node := root
		for _, r := range strconv.Itoa(i) {
			if node.mp[r] != nil {
				node = node.mp[r]
				cnt++
			} else {
				break
			}
		}
		if cnt > mx {
			mx = cnt
		}
	}
	return mx
}
