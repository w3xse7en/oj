package _022_07_24

import (
	"fmt"
	"strconv"
	"testing"
)

func equalPairs(grid [][]int) int {
	pairs := map[string]int{}
	list := []string{}
	for i := 0; i < len(grid); i++ {
		p := ""
		s := ""
		for j := 0; j < len(grid[i]); j++ {
			p += strconv.Itoa(grid[j][i]) + ","
			s += strconv.Itoa(grid[i][j]) + ","
		}
		list = append(list, s)
		pairs[p]++
	}
	cnt := 0
	for _, s := range list {
		cnt += pairs[s]
	}
	return cnt
}

func TestEqualPairs(t *testing.T) {
	//fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))
	//fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
	//fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 4}, {2, 4, 2, 2}, {2, 5, 2, 2}}))
	fmt.Println(equalPairs([][]int{{11, 1}, {1, 11}}))
}
