package maximumBags

import (
	"fmt"
	"sort"
	"testing"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	for i := range capacity {
		capacity[i] -= rocks[i]
	}
	sort.Ints(capacity)
	var cnt int
	for _, v := range capacity {
		additionalRocks -= v
		if additionalRocks >= 0 {
			cnt++
		}
	}
	return cnt
}

func TestB(t *testing.T) {
	fmt.Println(maximumBags([]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2))
}
