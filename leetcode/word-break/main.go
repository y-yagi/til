package main

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, str := range wordDict {
		dict[str] = true
	}

	wordStart := make([]bool, len(s)+1)
	wordStart[0] = true
	for i := 1; i < len(wordStart); i++ {
		for j := i - 1; j >= 0; j-- {
			if wordStart[j] && dict[s[j:i]] {
				wordStart[i] = true
			}
		}
	}
	return wordStart[len(s)]
}
