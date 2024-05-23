package _506_similarPairs

func similarPairs(words []string) int {
	same := map[string]int{}
	for _, word := range words {
		char := make([]byte, 26)
		for _, r := range word {
			char[r-'a'] = 1
		}
		same[string(char)]++
	}
	total := 0
	for _, cnt := range same {
		total += f(cnt)
	}
	return total
}

func f(v int) int {
	if v == 0 {
		return 0
	}
	if v == 1 {
		return 0
	}
	if v == 2 {
		return 1
	}
	return f(v-1) + v - 1
}
