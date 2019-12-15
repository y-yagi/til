package main

func searchInsert(nums []int, target int) int {
	start := len(nums) / 2
	middle := nums[start]
	answer := 0

	if target < middle {
		start = 0
	}

	for i := start; i < len(nums); i++ {
		if nums[i] >= target {
			answer = i
			break
		}
		answer = i + 1
	}

	return answer
}
