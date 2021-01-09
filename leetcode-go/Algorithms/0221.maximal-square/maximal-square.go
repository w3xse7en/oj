package problem0221

var maxX, maxY, minXY int

func maximalSquare(matrix [][]byte) int {
	var result int
	maxX, maxY = len(matrix), len(matrix[0])
	minXY = maxX
	if maxX > maxY {
		minXY = maxY
	}
	for x, raws := range matrix {
		for y, v := range raws {
			if v == '1' {
				maxSquare := findMaxSquare(matrix, x, y)
				if maxSquare > result {
					result = maxSquare
				}
			}
		}
	}
	return result * result
}

func findMaxSquare(matrix [][]byte, x, y int) int {
	var step int
	for ; step <= minXY; step++ {
		if !isNextStepOk(matrix, x, y, step) {
			return step
		}
	}
	return step
}

func isNextStepOk(matrix [][]byte, x, y, step int) bool {
	if x+step >= maxX || y+step >= maxY {
		return false
	}
	for i := 0; i < step; i++ {
		if matrix[x+i][y+step] == '0' {
			return false
		}
		if matrix[x+step][y+i] == '0' {
			return false
		}
	}
	if matrix[x+step][y+step] == '0' {
		return false
	}
	return true
}
