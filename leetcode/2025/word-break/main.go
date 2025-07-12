package main

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i < len(dp); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
