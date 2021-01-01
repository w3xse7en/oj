package problem0200

var (
	gridRow, gridCol int
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	gridCol = len(grid)
	gridRow = len(grid[0])
	cnt := 0
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == '1' {
				cnt++
				find(x, y, grid)
			}
		}
	}
	return cnt
}

func find(x, y int, grid [][]byte) {
	if x >= 0 && x < gridRow && y >= 0 && y < gridCol && grid[y][x] == '1' {
		grid[y][x] = '0'
		find(x+1, y, grid)
		find(x, y+1, grid)
		find(x-1, y, grid)
		find(x, y-1, grid)
	}
}
