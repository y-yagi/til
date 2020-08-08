package main

func removeElement(nums []int, val int) int {
	j := 0
	// Collect non-target values to the top.
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i] = nums[j]
			j++
		}
	}

	return j
}
