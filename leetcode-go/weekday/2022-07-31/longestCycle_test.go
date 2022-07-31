package _022_07_31

import (
	"fmt"
	"testing"
)

func TestLongestCycle(t *testing.T) {
	fmt.Println(longestCycle([]int{5, 8, -1, 5, -1, 6, 1, 6, 6, 5}))
	fmt.Println(longestCycle([]int{4, 3, 3, 4, 7, 2, 3, 3}))
	fmt.Println(longestCycle([]int{1, 5, 4, 2, 3, 6, 7, 8, 9, 0}))
	fmt.Println(longestCycle([]int{2, -1, 3, 1}))
}

func longestCycle(edges []int) int {
	maxDeep := -1
	for node, next := range edges {
		if next == -1 {
			continue
		}
		index := map[int]int{}
		deep := 0
		isLoop := false
		for {
			if node == -1 {
				break
			}
			if _, ok := index[node]; ok {
				isLoop = true
				deep = len(index) - index[node]
				break
			}
			index[node] = deep
			tmp := edges[node]
			edges[node] = -1
			node = tmp
			deep++
		}
		if isLoop && maxDeep < deep {
			maxDeep = deep
		}
	}
	return maxDeep
}
