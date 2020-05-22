package main

import (
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	a := make([]int, 26)
	for _, c := range tasks {
		a[c-'A']++
	}
	sort.Ints(a)
	maxVal := a[25] - 1
	idleSlots := maxVal * n
	for i := 24; i >= 0; i-- {
		idleSlots -= min(a[i], maxVal)
	}

	/*
		A - - - -
		A - - - -
		A - - - -
		A

		↓

		A B C D -
		A B C - -
		A B C - -
		A B
	*/

	if idleSlots > 0 {
		return idleSlots + len(tasks)
	}

	return len(tasks)
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
