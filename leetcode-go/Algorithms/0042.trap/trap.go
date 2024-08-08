package _042_trap

func trap1(height []int) int {
	list := []int{}
	sum := 0
	for i, v := range height {
		l := len(list)
		if v > l {
			list = append(list, make([]int, v-l)...)
		}
		for j := 0; j < v; j++ {
			if list[j] != 0 && i+1-list[j] >= 2 {
				sum += i + 1 - list[j] - 1
			}
			list[j] = i + 1
		}
	}
	return sum
}

func trap(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}
