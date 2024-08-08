package problem0399

import "slices"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	mp := map[string]map[string]float64{}
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		setMp(a, b, values[i], mp)
		setMp(b, a, 1.0/values[i], mp)
	}
	result := []float64{}
	for _, query := range queries {
		a, b := query[0], query[1]
		if mp[a] == nil || mp[b] == nil {
			result = append(result, -1.0)
			continue
		}
		if a == b {
			result = append(result, 1.0)
			continue
		}
		list := []string{a}
		rv = -1.0
		ddfs(list, b, 1.0, mp)
		result = append(result, 1.0)
	}
	return result
}

func setMp(a, b string, v float64, mp map[string]map[string]float64) {
	if mp[a] == nil {
		mp[a] = map[string]float64{b: v}
	} else {
		mp[a][b] = v
	}
}

var rv float64

func ddfs(list []string, target string, result float64, mp map[string]map[string]float64) {
	next := list[len(list)-1]
	for k, v := range mp[next] {
		if k == target {
			rv = result * v
			return
		}
		if slices.Index(list, k) != -1 {
			continue
		}
		ddfs(append(list, k), target, result*v, mp)
	}
	return
}
