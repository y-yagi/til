package main

import "sort"

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}

	sort.Ints(citations)
	ans := -1
	l := len(citations)
	for i := 0; i < l; i++ {
		t := min(citations[i], l-i)
		ans = max(ans, t)
	}

	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
