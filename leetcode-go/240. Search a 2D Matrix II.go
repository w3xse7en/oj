package main

import "fmt"

const (
	G = 1
	L = 2
	E = 3
)

func cp(src, tar int) int {
	if tar > src {
		return G
	} else if tar < src {
		return L
	} else {
		return E
	}
}

func searchMatrix(matrix [][]int, target int) bool {

	return true
}

func main() {
	m := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	matrix := searchMatrix(m, 4)
	fmt.Println(matrix)
}
