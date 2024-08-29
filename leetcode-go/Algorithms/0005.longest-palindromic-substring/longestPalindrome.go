package problem0005

func longestPalindrome(s string) string {
	mx := 0
	result := ""
	for i := range s {
		left, right := 0, 0
		l, r := i-1, i
		if l >= 0 && s[l] != s[r] {
			r = i + 1
		}
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				left = l
				right = r
				break
			}
			l--
			r++
		}
		if mx <= right+1-left {
			mx = right + 1 - left
			result = s[left : right+1]
		}
	}
	return result
}

func pLen(s string) int {
	l := len(s)
	for i, j := 0, l-1; i < l/2; i++ {
		if s[i] != s[j] {
			return 0
		}
		j--
	}
	return len(s)
}
