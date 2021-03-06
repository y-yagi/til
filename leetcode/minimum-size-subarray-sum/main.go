package main

import "math"

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	ans := math.MaxInt64
	left := 0
	sum := 0

	// Using two pointers.
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= s {
			ans = min(ans, i+1-left)
			sum -= nums[left]
			left++
		}
	}

	if ans != math.MaxInt64 {
		return ans
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
