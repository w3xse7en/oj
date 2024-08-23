package _134_canCompleteCircuit

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	idx, sum, total := -1, 0, 0
	for i := 0; i < n; i++ {
		sub := gas[i] - cost[i]
		total += sub
		sum += sub
		if sum < 0 {
			sum = 0
			idx = -1
		} else if sum >= 0 && idx == -1 {
			idx = i
		}
	}
	if idx < 0 || total < 0 {
		return -1
	}
	return idx
}
