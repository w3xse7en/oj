package problem0064

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for y := 0; y < m; y++ {
		if y == 0 {
			for x := range grid[y] {
				if x != 0 {
					grid[y][x] += grid[y][x-1]
				}
			}
		} else {
			for x := range grid[y] {
				if x == 0 {
					grid[y][x] += grid[y-1][x]
				} else {
					j := grid[y-1][x]
					i := grid[y][x-1]
					if grid[y][x]+i > grid[y][x]+j {
						grid[y][x] += j
					} else {
						grid[y][x] += i
					}
				}
			}
		}
	}
	return grid[m-1][n-1]
}
