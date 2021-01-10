package main

import (
	"sort"
)

func longestStrChain(words []string) int {
	dp := map[string]int{}
	ans := 0

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for _, w := range words {
		for i := 0; i < len(w); i++ {
			dp[w] = max(dp[w], dp[string(w[0:i])+string(w[i+1:])]+1)
			ans = max(ans, dp[w])
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
