package problem0121

func maxProfit(prices []int) int {
	min, max, result := 0xffffff, 0, 0
	for _, v := range prices {
		if min > v {
			min = v
			max = v
		}
		if max < v {
			max = v
			if max-min > result {
				result = max - min
			}
		}
	}
	return result
}
