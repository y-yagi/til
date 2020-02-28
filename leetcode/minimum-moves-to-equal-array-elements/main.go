package main

import "math"

func minMoves(nums []int) int {
	min := math.MaxInt32
	sum := 0
	l := len(nums)

	for _, v := range nums {
		sum += v
		if min > v {
			min = v
		}
	}

	return sum - min*l
}
