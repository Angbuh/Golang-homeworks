package wordcounter

import (
	"strings"
	"unicode"
)

func CountWords(text string) map[string]int {
	words := strings.Fields(text)
	m := make(map[string]int)
	for i := range words {
		words[i] = strings.TrimFunc(strings.ToLower(words[i]), unicode.IsPunct)
		words[i] = strings.ToLower(words[i])
	}
	for _, word := range words {
		if word != "" {
			m[word]++
		}
	}
	return m
}
