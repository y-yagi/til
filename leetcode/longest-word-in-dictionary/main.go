package main

import (
	"sort"
)

func longestWord(words []string) string {
	sort.Slice(words, func(i int, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] > words[j]
		}
		return len(words[i]) < len(words[j])
	})

	set := make(map[string]struct{})
	set[""] = struct{}{}

	var res string
	for _, word := range words {
		if _, found := set[word[:len(word)-1]]; found {
			set[word] = struct{}{}
			res = word
		}
	}
	return res
}
