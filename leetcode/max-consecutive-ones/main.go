package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 && (count == 0 || nums[i] == nums[i-1]) {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}

	if count > max {
		return count
	}
	return max
}
