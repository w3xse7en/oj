package problem0018

import (
	"sort"
)

type Eq struct {
	a, b, c, d int
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	a, b, c, d := 0, 1, len(nums)-2, len(nums)-1
	eqs := map[Eq]bool{}
	for a < d {
		for a < d {
			for ; b < c; b++ {
				na := nums[a]
				nb := nums[b]
				nc := nums[c]
				nd := nums[d]
				if na+nb+nc+nd == target {
					eq := Eq{
						a: na,
						b: nb,
						c: nc,
						d: nd,
					}
					eqs[eq] = true
				}
			}
			c--
			b = a + 1
			if b > c {
				d--
				b = a + 1
				c = d - 1
			}
		}
		a++
		d = len(nums) - 1
		b = a + 1
		c = d - 1
	}
	r := [][]int{}
	for eq := range eqs {
		r = append(r, []int{eq.a, eq.b, eq.c, eq.d})
	}

	return r
}
