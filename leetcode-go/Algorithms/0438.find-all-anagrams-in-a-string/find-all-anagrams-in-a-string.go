package problem0438

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	word := make([]int, 26)
	for k := range p {
		word[p[k]-'a']++
	}
	win := make([]int, 26)
	i, j := 0, 0
	for ; j < len(p); j++ {
		win[s[j]-'a']++
	}
	result := []int{}
	if match(win, word) {
		result = append(result, i)
	}
	for j < len(s) {
		win[s[i]-'a']--
		win[s[j]-'a']++
		i++
		j++
		if match(win, word) {
			result = append(result, i)
		}
	}
	return result
}

func match(win, word []int) bool {
	for idx := range word {
		if word[idx] != win[idx] {
			return false
		}
	}
	return true
}
