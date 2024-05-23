package problem0073

func setZeroes(matrix [][]int) {
	zi, zj := false, false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				if i == 0 {
					zj = true
				}
				if j == 0 {
					zi = true
				}
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
		if zi {
			matrix[i][0] = 0
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
		if zj {
			matrix[0][j] = 0
		}
	}
	if zi || zj {
		matrix[0][0] = 0
	}
	return
}
