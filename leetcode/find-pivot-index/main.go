package main

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	var sum, leftsum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if leftsum == sum-leftsum-nums[i] {
			return i
		}
		leftsum += nums[i]
	}

	return -1
}
