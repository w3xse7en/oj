package _289_gameOfLife

import "fmt"

func gameOfLife(board [][]int) {
	n := [][]int{}
	for i := 0; i < len(board); i++ {
		t := []int{}
		for j := 0; j < len(board[0]); j++ {
			t = append(t, f(i, j, board))
		}
		n = append(n, t)
	}
	copy(board, n)
}

func f(i, j int, board [][]int) int {
	m := len(board)
	n := len(board[0])
	l, d := 0, 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			a, b := i+x, j+y
			if a == i && b == j {
				continue
			}
			if a < 0 || a >= m {
				continue
			}
			if b < 0 || b >= n {
				continue
			}
			if board[a][b] == 0 {
				d += 1
			}
			if board[a][b] == 1 {
				l += 1
			}
		}
	}
	fmt.Println("i", i, "j", j, "l", l, "d", d)
	if board[i][j] == 1 {
		if l < 2 {
			return 0
		}
		if l > 3 {
			return 0
		}
	}
	if board[i][j] == 0 {
		if l == 3 {
			return 1
		}
	}
	return board[i][j]
}
