package _134_canCompleteCircuit

func canCompleteCircuit(gas []int, cost []int) int {
	i := 0
	j := -1
	remain := gas[0]
	for cnt := 0; cnt < len(gas); cnt++ {
		ni := newIdx(i+1, len(gas))
		if remain-cost[i] < 0 {
			j = ni
			i = ni
			cnt = 0
			remain = gas[ni]
		} else {
			remain = remain - cost[i] + gas[ni]
		}
	}
	if i != j {
		return -1
	}
	return i
}

func newIdx(idx, length int) int {
	if idx >= length {
		return 0
	}
	return idx
}
