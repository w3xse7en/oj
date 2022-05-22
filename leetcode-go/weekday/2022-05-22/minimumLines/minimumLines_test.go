package minimumLines

import (
	"fmt"
	"sort"
	"testing"
)

func minimumLines(stockPrices [][]int) int {
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] > stockPrices[j][0]
	})
	var cnt int
	d, p := 0, 0
	for i := 1; i < len(stockPrices); i++ {
		sd, sp := stockPrices[i][0]-stockPrices[i-1][0], stockPrices[i][1]-stockPrices[i-1][1]
		if i == 1 {
			cnt = 1
		} else if d*sp != sd*p {
			cnt++
		}
		d, p = sd, sp
	}
	return cnt
}

func TestMinimumLines(t *testing.T) {
	fmt.Println(minimumLines([][]int{
		{72, 98},
		{62, 27},
		{32, 7},
		{71, 4},
		{25, 19},
		{91, 30},
		{52, 73},
		{10, 9},
		{99, 71},
		{47, 22},
		{19, 30},
		{80, 63},
		{18, 15},
		{48, 17},
		{77, 16},
		{46, 27},
		{66, 87},
		{55, 84},
		{65, 38},
		{30, 9},
		{50, 42},
		{100, 60},
		{75, 73},
		{98, 53},
		{22, 80},
		{41, 61},
		{37, 47},
		{95, 8},
		{51, 81},
		{78, 79},
		{57, 95},
	}))
}
