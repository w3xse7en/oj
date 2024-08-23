package _068_fullJustify

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	lineWords := []string{}
	lineLength := 0
	result := []string{}
	for i := 0; i < len(words); i++ {
		word := words[i]
		wordLen := len(word)
		lineWords = append(lineWords, word)
		if lineLength+wordLen+len(lineWords)-1 <= maxWidth {
			lineLength += wordLen
			continue
		}
		result = append(result, getLineWords(lineWords[:len(lineWords)-1], lineLength, maxWidth, false))
		lineLength = len(word)
		lineWords = lineWords[len(lineWords)-1:]
	}
	result = append(result, getLineWords(lineWords, lineLength, maxWidth, true))
	return result
}

func getLineWords(words []string, wordsLen, maxLen int, isEnd bool) string {
	if len(words) == 0 {
		return ""
	}
	space := buildSpaceForWords(words, wordsLen, maxLen, isEnd)
	var result strings.Builder
	for i, word := range words {
		result.WriteString(word)
		if i < len(space) {
			result.WriteString(space[i])
		}
	}
	return result.String()
}

func buildSpaceForWords(words []string, wordsLen, maxLen int, isEnd bool) []string {
	space := make([]string, len(words)-1)
	spaceTotalLen := maxLen - wordsLen
	//
	if isEnd {
		for i := range space {
			space[i] = " "
		}
		space = append(space, strings.Repeat(" ", spaceTotalLen-len(space)))
		return space
	}
	//
	if len(space) == 0 {
		space = append(space, strings.Repeat(" ", spaceTotalLen))
		return space
	}
	//
	s := spaceTotalLen % len(space)
	sp := spaceTotalLen / len(space)
	for i := range space {
		sn := sp
		if s > 0 {
			sn += 1
		}
		space[i] = strings.Repeat(" ", sn)
		s--
	}
	return space
}
