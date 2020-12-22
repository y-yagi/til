package main

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	l := len(nums) - 1
	return max(nums[0]*nums[1]*nums[l], nums[l]*nums[l-1]*nums[l-2])
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
