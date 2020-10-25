package main

import "sort"

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	hi := len(tokens) - 1
	var lo, points, ans int
	for lo <= hi && (P >= tokens[lo] || points > 0) {
		for lo <= hi && P >= tokens[lo] {
			P -= tokens[lo]
			lo++
			points++
		}

		ans = max(ans, points)
		if lo <= hi && points > 0 {
			P += tokens[hi]
			hi--
			points--
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
