package problem0011

func maxArea(height []int) int {
	tmp, max := 0, 0
	for xi, xj := 0, len(height)-1; xi < xj; {
		yi, yj := height[xi], height[xj]
		x := xj - xi
		if yi > yj {
			tmp = x * yj
			xj--
		} else {
			tmp = x * yi
			xi++
		}
		if tmp > max {
			max = tmp
		}
	}

	return max
}
