package main

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	answer := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(answer) == 0 || answer[len(answer)-1][1] < intervals[i][1] {
			answer = append(answer, intervals[i])
		}
	}
	return len(answer)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
