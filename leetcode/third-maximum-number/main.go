package main

import (
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	l := len(nums)

	if l < 3 {
		return nums[l-1]
	}

	before := nums[l-1]
	pos := 2
	for i := l - 2; i > -1; i-- {
		if before != nums[i] {
			before = nums[i]
			pos--
		}

		if pos == 0 {
			return nums[i]
		}

	}
	return nums[l-1]
}
