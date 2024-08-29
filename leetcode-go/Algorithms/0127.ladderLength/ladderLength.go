package _127_ladderLength

import (
	"math"
	"slices"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	ei := slices.Index(wordList, endWord)
	if ei == -1 {
		return 0
	}
	if beginWord == endWord || isOneDiff(beginWord, endWord) {
		return 2
	}
	bl := []string{beginWord}
	el := []string{endWord}
	bvis := map[string]int{beginWord: 1}
	evis := map[string]int{endWord: 1}
	for i := 2; i < 1000; i++ {
		bl, bvis = nextList(bl, wordList, bvis, i)
		el, evis = nextList(el, wordList, evis, i)
		if len(bl) == 0 || len(el) == 0 {
			return 0
		}
		has := false
		mi := math.MaxInt
		for w, d := range bvis {
			if evis[w] != 0 {
				mi = min(mi, d+evis[w]-1)
			}
		}
		if has {
			return mi
		}
	}
	return 0
}
func isOneDiff(word, target string) bool {
	if len(word) != len(target) || word == target {
		return false
	}
	diffCount := 0
	for i := range word {
		if target[i] != word[i] {
			diffCount++
		}
		if diffCount > 1 {
			return false
		}
	}
	return true
}

func nextList(list, wordList []string, vis map[string]int, deep int) ([]string, map[string]int) {
	next := []string{}
	for _, w := range list {
		for _, v := range wordList {
			if vis[v] == 0 && isOneDiff(w, v) {
				next = append(next, v)
				vis[v] = deep
			}
		}
	}
	return next, vis
}
