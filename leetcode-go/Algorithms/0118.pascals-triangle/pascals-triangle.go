package problem0118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	result := [][]int{{1}}
	for i := 2; i <= numRows; i++ {
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
	return result
}
