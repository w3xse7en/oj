package _139_wordBreak

var cache map[int]bool

func wordBreak(s string, wordDict []string) bool {
	mx := 0
	cache = map[int]bool{}
	word := map[string]bool{}
	for _, v := range wordDict {
		word[v] = true
		mx = max(mx, len(v))
	}
	return dfs(0, mx, s, word)
}

func dfs(idx, mx int, str string, exist map[string]bool) bool {
	if idx >= len(str) {
		return true
	}
	if v, ok := cache[idx]; ok {
		return v
	}
	for i := idx + 1; i <= len(str); i++ {
		if exist[str[idx:i]] && dfs(i, mx, str, exist) {
			cache[i] = true
			return true
		}
	}
	cache[idx] = false
	return false
}
