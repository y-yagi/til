package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	answer := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(answer) == 0 || answer[len(answer)-1][1] < intervals[i][0] {
			answer = append(answer, intervals[i])
		} else {
			answer[len(answer)-1][1] = max(answer[len(answer)-1][1], intervals[i][1])
		}
	}
	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
