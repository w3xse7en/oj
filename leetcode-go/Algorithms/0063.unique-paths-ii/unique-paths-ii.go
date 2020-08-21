package problem0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	jStop := false
	for i, j := 0, 0; j < n; j++ {
		if obstacleGrid[i][j] == 1 {
			obstacleGrid[i][j] = 0
			jStop = true
		}
		if !jStop {
			obstacleGrid[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			if j == 0 {
				obstacleGrid[i][j] = 0 + obstacleGrid[i-1][j]
			} else {
				obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}
