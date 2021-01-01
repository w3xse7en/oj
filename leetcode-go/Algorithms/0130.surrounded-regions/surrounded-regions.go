package problem0130

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	x, y := 0, 0
	row, col := len(board[0]), len(board)

	f := func(x, y int, board [][]byte) {
		if board[y][x] == 'O' {
			fill(x, y, board, 'O', 'A')
		}
		if board[y][x] == 'X' {
			fill(x, y, board, 'X', 'B')
		}
	}
	for ; x < row; x++ {
		f(x, y, board)
	}
	x--
	for ; y < col; y++ {
		f(x, y, board)
	}
	y--
	for ; x >= 0; x-- {
		f(x, y, board)
	}
	x++
	for ; y >= 0; y-- {
		f(x, y, board)
	}
	y++

	for y, row := range board {
		for x := range row {
			if board[y][x] == 'A' {
				board[y][x] = 'O'
			} else if board[y][x] == 'B' {
				board[y][x] = 'X'
			} else if board[y][x] == 'O' {
				board[y][x] = 'X'
			}
		}
	}
	return
}

func fill(x, y int, board [][]byte, from, to byte) {
	if y >= 0 && y < len(board) && x >= 0 && x < len(board[0]) && board[y][x] == from {
		board[y][x] = to
		fill(x+1, y, board, from, to)
		fill(x, y+1, board, from, to)
		fill(x-1, y, board, from, to)
		fill(x, y-1, board, from, to)
	}
}
