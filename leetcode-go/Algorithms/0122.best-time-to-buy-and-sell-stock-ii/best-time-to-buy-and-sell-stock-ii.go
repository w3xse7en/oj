package problem0122

func maxProfit(prices []int) int {
	result := 0
	for i := 0; i < len(prices); i++ {
		tmp := prices[i]
		jump := i
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > tmp {
				tmp = prices[j]
				jump = j
			} else {
				jump = j - 1
				break
			}
		}
		result += tmp - prices[i]
		i = jump
	}
	return result
}

func maxProfitBetter(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	min := prices[0]

	for _, price := range prices {
		if price > min {
			profit += price - min
		}
		min = price
	}

	return profit
}
