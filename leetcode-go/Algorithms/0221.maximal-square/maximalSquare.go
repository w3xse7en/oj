package problem0221

func maximalSquare(matrix [][]byte) int {
	m := 0
	dp := make([][]int, len(matrix))
	for i, bytes := range matrix {
		dp[i] = make([]int, len(bytes))
		for j := range bytes {
			dp[i][j] = int(matrix[i][j] - '0')
			if i-1 >= 0 && j-1 >= 0 && matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
			m = max(m, dp[i][j])
		}
	}
	return m * m
}

func maximalSquare3(matrix [][]byte) int {
	m := 0
	for i, bytes := range matrix {
		for j := range bytes {
			if matrix[i][j] == '1' {
				m = max(maximalS(i, j, matrix), m)
			}
		}
	}
	return m * m
}

func maximalS(a, b int, matrix [][]byte) int {
	am := len(matrix)
	bm := len(matrix[0])
	mi := min(am-a, bm-b)
	//fmt.Println("a,b", a, b)
	m := 1
	for ; m < mi; m++ {
		has := false
		//fmt.Println(m, mi)
		for i := 0; i <= m; i++ {
			//fmt.Println(a+m, b+i, "|", a+i, b+m)
			if matrix[a+m][b+i] == '0' {
				has = true
				break
			}
			if matrix[a+i][b+m] == '0' {
				has = true
				break
			}
		}
		if has {
			return m
		}
	}
	return m
}
