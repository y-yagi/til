package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	diff := math.MaxInt32
	sz := len(nums)
	sort.Ints(nums)
	for i := 0; i < sz && diff != 0; i++ {
		lo := i + 1
		hi := sz - 1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			if abs(target-sum) < abs(diff) {
				diff = target - sum
			}
			if sum < target {
				lo++
			} else {
				hi--
			}
		}
	}
	return target - diff
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
