package _135_candy

func candy(ratings []int) int {
	list := make([]int, len(ratings))
	list[0] = 1
	for i := 1; i < len(ratings); i++ {
		list[i] = 1
		ra, rb := ratings[i-1], ratings[i]
		if ra < rb {
			list[i] = list[i-1] + 1
		}
	}
	sum := list[len(list)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		ra, rb := ratings[i], ratings[i+1]
		if ra > rb {
			list[i] = max(list[i], list[i+1]+1)
		}
		sum += list[i]
	}
	return sum
}
