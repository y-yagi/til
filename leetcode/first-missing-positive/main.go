package main

import "sort"

func firstMissingPositive(nums []int) int {
	smaller := 1
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == smaller {
			smaller++
		}
	}

	return smaller
}
