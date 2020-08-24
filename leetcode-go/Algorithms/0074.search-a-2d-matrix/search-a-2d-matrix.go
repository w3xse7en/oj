package problem0074

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return searchCol(matrix, target)
}
func searchCol(col [][]int, target int) bool {
	i := 0
	j := len(col) - 1
	mid := (i + j) / 2
	for {
		if i > j {
			if i-1 < 0 {
				return false
			}
			return searchRow(col[i-1], target)
		}
		if len(col[mid]) == 0 {
			return false
		}
		if col[mid][0] == target {
			return true
		} else if col[mid][0] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
		mid = (i + j) / 2
	}
}
func searchRow(row []int, target int) bool {
	i := 0
	j := len(row) - 1
	mid := (i + j) / 2
	for {
		if i > j {
			return false
		}
		if row[mid] == target {
			return true
		} else if row[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
		mid = (i + j) / 2
	}
}
