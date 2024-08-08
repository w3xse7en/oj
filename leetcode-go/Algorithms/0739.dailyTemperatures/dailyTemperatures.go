package _739_dailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	type ti struct {
		t   int
		idx int
	}
	list := make([]ti, 0, len(temperatures))
	result := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		cut := -1
		for j := len(list) - 1; j >= 0; j-- {
			v := list[j]
			if temperature > v.t {
				result[v.idx] = i - v.idx
				cut = j
			} else {
				break
			}
		}
		if cut >= 0 {
			list = append(list[:cut])
		}
		list = append(list, ti{idx: i, t: temperature})
	}
	return result
}
