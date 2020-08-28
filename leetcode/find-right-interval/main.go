package main

import "math"

func findRightInterval(intervals [][]int) []int {
	ans := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		min := math.MaxInt32
		minindex := -1
		for j := 0; j < len(intervals); j++ {
			if intervals[j][0] >= intervals[i][1] && intervals[j][0] < min {
				min = intervals[j][0]
				minindex = j
			}
			ans[i] = minindex
		}
	}

	return ans
}

func findRightInterval_withMap(intervals [][]int) []int {
	ans := []int{}
	dict := map[int]int{}

	for i := 0; i < len(intervals); i++ {
		dict[intervals[i][0]] = i
	}

	for i := 0; i < len(intervals); i++ {
		if v, found := dict[intervals[i][1]]; found {
			ans = append(ans, v)
			continue
		}

		value := -1
		for k, v := range dict {
			if k > intervals[i][1] {
				if value == -1 {
					value = v
				} else {
					value = min(value, v)
				}
			}
		}

		ans = append(ans, value)
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
