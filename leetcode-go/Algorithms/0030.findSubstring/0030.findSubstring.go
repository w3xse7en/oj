package _030_findSubstring

import "math"

func findSubstring(s string, words []string) []int {
	minLen, maxLen := math.MaxInt, math.MinInt
	totalWordCnt, totalWordLen := 0, 0
	wordMap := map[string]int{}
	wordCnt := []int{0}
	charMap := map[uint8]bool{}
	for _, word := range words {
		idx, ok := wordMap[word]
		if ok {
			wordCnt[idx]++
		} else {
			wordMap[word] = len(wordCnt)
			wordCnt = append(wordCnt, 1)
		}
		charMap[word[0]] = true
		totalWordCnt++
		totalWordLen += len(word)
		minLen = min(minLen, len(word))
		maxLen = max(maxLen, len(word))
	}
	idx := []int{}
	for i := 0; i < len(s); i++ {
		j := i + totalWordLen
		if j <= len(s) && charMap[s[i]] {
			tmpCnt := make([]int, len(wordCnt))
			copy(tmpCnt, wordCnt)
			if dfs(s, i, j, wordMap, tmpCnt, totalWordCnt, minLen, maxLen) {
				idx = append(idx, i)
			}
		}
	}
	return idx
}

func dfs(s string, st, ed int, wordMap map[string]int, wordCnt []int, totalWordCnt, minlen, maxlen int) bool {
	if totalWordCnt == 0 && st == ed {
		return true
	}
	for i := minlen; i <= maxlen && i <= len(s); i++ {
		idx := wordMap[s[st:st+i]]
		if wordCnt[idx] > 0 {
			wordCnt[idx]--
			totalWordCnt--
			return dfs(s, st+i, ed, wordMap, wordCnt, totalWordCnt, minlen, maxlen)
		}
	}
	return false
}
