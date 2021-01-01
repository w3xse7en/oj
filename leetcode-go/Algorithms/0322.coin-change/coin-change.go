package problem0322

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	m := 0xffffffff
	for _, coin := range coins {
		v := coinChange(coins, amount-coin)
		if v == -1 {
			continue
		}
		if v+1 < m {
			m = v + 1
		}
	}
	return m
}
