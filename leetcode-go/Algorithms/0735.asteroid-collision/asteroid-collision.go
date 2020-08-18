package problem0735

func asteroidCollision(asteroids []int) []int {
	var r []int
	for _, v := range asteroids {
		r = f(r, v)
	}
	return r
}

func f(r []int, v int) []int {
	if len(r) == 0 || v > 0 {
		return append(r, v)
	}
	absV := -v
	lastIndex := len(r) - 1
	if r[lastIndex] < 0 {
		return append(r, v)
	}
	if r[lastIndex] == absV {
		return r[:lastIndex]
	}
	if r[lastIndex] > absV {
		return r
	}
	if r[lastIndex] < absV {
		return f(r[:lastIndex], v)
	}
	return r
}
