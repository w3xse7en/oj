package problem0119

func getRow(rowIndex int) []int {
	result := []int{1}
	for i := 1; i <= rowIndex; i++ {
		result = append(result, result[i-1]*(rowIndex-(i-1))/i)
	}
	return result
}

func getRowOrigin(rowIndex int) []int {

	result := [][]int{{1}}
	if rowIndex == 0 {
		return result[0]
	}
	for i := 2; i <= rowIndex; i++ {
		line := make([]int, i)
		for k := range line {
			i1, i2 := k-1, k
			if i1 == -1 {
				line[k] = result[i-2][i2]
			} else if i2 == i-1 {
				line[k] = result[i-2][i1]
			} else {
				line[k] = result[i-2][i1] + result[i-2][i2]
			}
		}
		result = append(result, line)
	}
	return result[rowIndex]
}
