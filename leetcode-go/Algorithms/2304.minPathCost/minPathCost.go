package _304_minPathCost

import (
	"fmt"
	"math"
	"slices"
)

func minPathCost(grid [][]int, moveCost [][]int) int {
	cost := grid[0]
	tmp := make([]int, len(grid[0]))
	for i := 1; i < len(grid); i++ {
		fmt.Println("1", tmp, cost)
		for j, v := range grid[i] {
			ct := math.MaxInt
			for k, c := range cost {
				ct = min(ct, c+moveCost[grid[i-1][k]][j]+v)
				fmt.Println("ct:", ct, " c:", c, " movec:", moveCost[grid[i-1][k]][j], " v", v)
			}
			tmp[j] = ct
		}
		fmt.Println()
		fmt.Println("2", tmp, cost)
		cost = tmp
	}
	return slices.Min(cost)
}

//1 [0 0 0 0 0 0 0 0 0] [28 35 29 5 13 17 18 23 14]
//2 [36 37 55 61 30 48 55 38 40] [28 35 29 5 13 17 18 23 14]
//1 [36 37 55 61 30 48 55 38 40] [36 37 55 61 30 48 55 38 40]
//2 [41 83 86 52 56 67 129 117 72] [41 83 86 52 56 67 129 117 72]
//1 [41 83 86 52 56 67 129 117 72] [41 83 86 52 56 67 129 117 72]
//2 [86 107 75 92 124 87 124 113 115] [86 107 75 92 124 87 124 113 115]

//1 [0 0 0 0 0 0 0 0 0] [28 35 29 5 13 17 18 23 14]
//2 [36 37 55 61 30 48 55 38 40] [28 35 29 5 13 17 18 23 14]
//1 [0 0 0 0 0 0 0 0 0] [36 37 55 61 30 48 55 38 40]
//2 [41 83 86 52 56 50 114 87 60] [36 37 55 61 30 48 55 38 40]
//1 [0 0 0 0 0 0 0 0 0] [41 83 86 52 56 50 114 87 60]
//2 [86 90 75 92 107 59 66 101 103] [41 83 86 52 56 50 114 87 60]
