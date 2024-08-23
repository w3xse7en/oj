package _985_sumEvenAfterQueries

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	sum := 0
	for _, num := range nums {
		if num%2 == 0 {
			sum += num
		}
	}
	list := make([]int, 0, len(queries))
	for _, query := range queries {
		val, idx := query[0], query[1]
		oriV := nums[idx]
		if oriV%2 == 0 {
			sum -= oriV
		}
		newV := oriV + val
		nums[idx] = newV
		if newV%2 == 0 {
			sum += newV
		}
		//fmt.Println(val, idx, oriV, newV, sum)
		list = append(list, sum)
	}
	return list
}
