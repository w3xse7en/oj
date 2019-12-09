package problem0054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	x, y := 0, 0
	xs, ys, xe, ye := 0, 0, len(matrix[0])-1, len(matrix)-1
	out := make([]int, 0)
	for xs <= xe && ys <= ye {
		ok := false
		for ; x <= xe; x++ {
			out = append(out, matrix[y][x])
			ok = true
		}
		if ok {
			x--
			y++
		} else {
			return out
		}
		ok = false
		for ; y <= ye; y++ {
			out = append(out, matrix[y][x])
			ok = true
		}
		if ok {
			x--
			y--
		} else {
			return out
		}
		ok = false
		for ; x >= xs; x-- {
			out = append(out, matrix[y][x])
			ok = true
		}
		if ok {
			x++
			y--
		} else {
			return out
		}
		ok = false
		for ; y > ys; y-- {
			out = append(out, matrix[y][x])
			ok = true
		}
		if ok {
			x++
			y++
		} else {
			return out
		}
		ye--
		xe--
		ys++
		xs++
	}
	return out
}
