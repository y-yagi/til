package main

import (
	"sort"
)

func heightChecker(heights []int) int {
	count := 0
	dup := make([]int, len(heights))
	copy(dup, heights)
	sort.Ints(dup)

	for i := 0; i < len(heights); i++ {
		if heights[i] != dup[i] {
			count++
		}
	}

	return count
}
