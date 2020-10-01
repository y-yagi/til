package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	contains := 0

	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			contains++
			break
		}
	}

	if contains == 0 {
		return 1
	}

	if n == 1 {
		return 2
	}

	// Replace negative numbers, zeros,
	// and numbers larger than n by 1s.
	// After this convertion nums will contain
	// only positive numbers.
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	// Use index as a hash key and number sign as a presence detector.
	// For example, if nums[1] is negative that means that number `1`
	// is present in the array.
	// If nums[2] is positive - number 2 is missing.
	for i := 0; i < n; i++ {
		a := abs(nums[i])
		if a == n {
			nums[0] = -abs(nums[0])
		} else {
			nums[a] = -abs(nums[a])
		}
	}

	// Now the index of the first positive number
	// is equal to first missing positive.
	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return n
	}

	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}
