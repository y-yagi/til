package main

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	firstEnd := points[0][1]
	for _, p := range points {
		xStart := p[0]
		xEnd := p[1]
		if firstEnd < xStart {
			arrows++
			firstEnd = xEnd
		}
	}

	return arrows
}
