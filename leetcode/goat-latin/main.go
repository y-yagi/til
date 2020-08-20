package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {
	vowels := map[byte]bool{'a': false, 'e': false, 'i': false, 'o': false, 'u': false, 'A': false, 'E': false, 'I': false, 'O': false, 'U': false}
	ans := []string{}
	words := strings.Split(S, " ")
	for i, word := range words {
		if _, found := vowels[word[0]]; found {
			ans = append(ans, fmt.Sprintf("%s%s%s", word, "ma", strings.Repeat("a", i+1)))
		} else {
			ans = append(ans, fmt.Sprintf("%s%c%s%s", word[1:], word[0], "ma", strings.Repeat("a", i+1)))
		}
	}

	return strings.Join(ans, " ")
}
