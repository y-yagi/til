package main

import (
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	queue := []int{intervals[0][1]}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= queue[0] {
			queue = queue[1:]
		}
		queue = append(queue, intervals[i][1])
		sort.Ints(queue)
	}

	return len(queue)
}
