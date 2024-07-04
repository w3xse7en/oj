package _400_canConstruct

func canConstruct(s string, k int) bool {
	h := make([]int, 26)
	for _, v := range s {
		h[v-'a']++
	}
	a := 0
	for _, cnt := range h {
		if cnt%2 != 0 {
			a++
		}
	}
	if k >= a && k <= len(s) {
		return true
	}
	return false
}
