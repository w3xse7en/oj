package problem0079

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			var result bool
			dsf(0, i, j, board, word, &result)
			if result {
				return true
			}
		}
	}
	return false
}

func dsf(idx, i, j int, board [][]byte, word string, result *bool) {
	if i >= 0 && j >= 0 && i <= len(board)-1 && j <= len(board[0])-1 {
		tmp := board[i][j]
		if idx <= len(word)-1 && word[idx] == board[i][j] {
			//fmt.Println(i, j, string(board[i][j]), idx, string(word[idx]))
			if idx == len(word)-1 {
				*result = true
				return
			}
			board[i][j] = 0
			dsf(idx+1, i, j+1, board, word, result)
			dsf(idx+1, i+1, j, board, word, result)
			dsf(idx+1, i, j-1, board, word, result)
			dsf(idx+1, i-1, j, board, word, result)
		}
		board[i][j] = tmp
	}
	return
}
