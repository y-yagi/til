package main

import "math"

func increasingTriplet(nums []int) bool {
	small := math.MaxInt64
	big := math.MaxInt64

	for _, n := range nums {
		if n <= small {
			small = n
		} else if n <= big {
			big = n
		} else {
			return true
		}
	}

	return false
}
