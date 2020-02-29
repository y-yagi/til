package main

import "sort"

func findContentChildren(g []int, s []int) int {
	answer := 0
	sort.Ints(g)
	sort.Ints(s)
	for i := 0; answer < len(g) && i < len(s); i++ {
		if g[answer] <= s[i] {
			answer++
		}
	}

	return answer
}
