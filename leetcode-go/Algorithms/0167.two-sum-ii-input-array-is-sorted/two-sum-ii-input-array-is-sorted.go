package problem0167

func twoSum(numbers []int, target int) []int {
	for i, number := range numbers {
		j := half(numbers[i+1:], target-number)
		if j != -1 {
			return []int{i + 1, j + i + 1 + 1}
		}
	}
	return nil
}

func half(numbers []int, target int) int {
	left, right := 0, len(numbers)-1
	for {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		numb := numbers[mid]
		if target == numb {
			return mid
		}
		if target > numb {
			left = mid + 1
		}
		if target < numb {
			right = mid - 1
		}
	}
}

func twoSumFast(numbers []int, target int) []int {
	mp := make(map[int]int, len(numbers))
	for i, number := range numbers {
		mp[number] = i + 1
	}
	for i, number := range numbers {
		j := mp[target-number]
		if j != 0 && i+1 < j {
			return []int{i + 1, j}
		}
	}
	return nil
}
