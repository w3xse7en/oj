package problem0059

func generateMatrix(n int) [][]int {
	ma := make([][]int, n)
	for i := 0; i < n; i++ {
		ma[i] = make([]int, n)
	}
	x, y := 0, 0
	option, i := 1, 0
	for i != n*n {
		if option == 1 {
			if x < n && ma[y][x] == 0 {
				i++
				ma[y][x] = i
				x++
			} else {
				option = 2
				x--
				y++
			}
		}
		if option == 2 {
			if y < n && ma[y][x] == 0 {
				i++
				ma[y][x] = i
				y++
			} else {
				option = 3
				y--
				x--
			}
		}
		if option == 3 {
			if x >= 0 && ma[y][x] == 0 {
				i++
				ma[y][x] = i
				x--
			} else {
				option = 4
				x++
				y--
			}
		}
		if option == 4 {
			if y >= 0 && ma[y][x] == 0 {
				i++
				ma[y][x] = i
				y--
			} else {
				option = 1
				y++
				x++
			}
		}
	}
	return ma
}
