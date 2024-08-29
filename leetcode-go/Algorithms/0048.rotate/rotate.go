package _048_rotate

func rotate(matrix [][]int) {
	l := len(matrix)
	for step := 0; step < l/2; step++ {
		i1, j1 := step, step
		i2, j2 := step, l-step-1
		i3, j3 := l-step-1, l-step-1
		i4, j4 := l-step-1, step
		for i := 0; i < l-step*2-1; i++ {
			matrix[i1][j1], matrix[i2][j2], matrix[i3][j3], matrix[i4][j4] = matrix[i4][j4], matrix[i1][j1], matrix[i2][j2], matrix[i3][j3]
			j1++
			i2++
			j3--
			i4--
		}
	}
}
