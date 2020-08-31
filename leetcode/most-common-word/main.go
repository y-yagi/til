package main

import (
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)

	paragraphs := map[rune]bool{'!': false, '?': false, '\'': false, ',': false, ';': false, '.': false, ' ': false}
	words := []string{}
	t := []rune{}

	for _, b := range paragraph {
		if _, found := paragraphs[b]; found {
			if len(t) != 0 {
				words = append(words, string(t))
				t = []rune{}
			}
			continue
		}
		t = append(t, b)
	}

	if len(t) != 0 {
		words = append(words, string(t))
	}

	counter := map[string]int{}
	bannedMap := map[string]bool{}

	for _, word := range banned {
		bannedMap[word] = true
	}

	for _, word := range words {
		if _, found := bannedMap[word]; found {
			continue
		}
		counter[word] += 1
	}

	ans := ""
	maxCount := 0
	for k, v := range counter {
		if v > maxCount {
			ans = k
			maxCount = v
		}
	}

	return ans
}
