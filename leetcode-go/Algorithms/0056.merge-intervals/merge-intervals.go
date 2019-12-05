package problem0056

import (
	"sort"
)

type SortInt struct {
	Min, Max int
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sortMap := make(map[*SortInt]int)
	sortInt := make([]*SortInt, 0)
	for k, v := range intervals {
		t := &SortInt{
			Min: v[0],
			Max: v[1],
		}
		sortMap[t] = k
		sortInt = append(sortInt, t)
	}
	sort.Slice(sortInt, func(i, j int) bool {
		if sortInt[i].Min < sortInt[j].Min {
			return true
		} else if sortInt[i].Min == sortInt[j].Min && sortInt[i].Max < sortInt[j].Max {
			return true
		}
		return false
	})

	result := [][]int{}
	min, max := intervals[sortMap[sortInt[0]]][0], intervals[sortMap[sortInt[0]]][1]
	if min > max {
		min, max = max, min
	}
	for i := 1; i < len(sortInt); i++ {
		nextMin := intervals[sortMap[sortInt[i]]][0]
		nextMax := intervals[sortMap[sortInt[i]]][1]
		if nextMin > nextMax {
			nextMin, nextMax = nextMax, nextMin
		}
		if nextMax < max {
			min, max, nextMax, nextMin = nextMin, nextMax, max, min
		}
		if max < nextMin {
			result = append(result, []int{min, max})
			min = nextMin
			max = nextMax
		} else {
			if nextMin < min {
				min = nextMin
			}
			if nextMax > max {
				max = nextMax
			}
		}
	}
	result = append(result, []int{min, max})
	return result
}
