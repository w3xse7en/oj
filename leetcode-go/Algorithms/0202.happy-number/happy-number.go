package problem0202

func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = digit(slow)
		fast = digit(digit(fast))
		if slow == 1 || fast == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
}

func digit(n int) int {
	sum := 0
	for n > 0 {
		s := n % 10
		sum += s * s
		n /= 10
	}
	return sum
}

func own(n int) bool {
	exist := map[int]bool{n: true}
	for {
		var ns []int
		for n > 0 {
			ns = append(ns, n%10)
			n /= 10
		}
		for _, v := range ns {
			n += v * v
		}
		if n == 1 {
			return true
		}
		if exist[n] {
			return false
		} else {
			exist[n] = true
		}
	}
}
