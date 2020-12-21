package main

func minStartValue(nums []int) int {
	sum := 0
	minv := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		minv = min(minv, sum)
	}

	return 1 - minv
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
