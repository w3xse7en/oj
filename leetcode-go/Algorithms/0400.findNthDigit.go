package Algorithms

func findNthDigit(n int) int {
	var list = []int{1}
	for idx := 1; idx <= n; {
		for i := len(list) - 1; i >= 0; i-- {
			if idx == n {
				return list[i]
			}
			idx++
		}
		list = add(list)
	}
	return 0
}
func add(list []int) []int {
	list[0]++
	for i := range list {
		if list[i] < 10 {
			break
		}
		list[i] = 0
		if i+1 < len(list) {
			list[i+1]++
		} else {
			list = append(list, 1)
		}
	}
	return list
}
