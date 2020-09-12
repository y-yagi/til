package main

func maxProduct(nums []int) int {
	maxSoFar := nums[0]
	minSoFar := nums[0]
	result := maxSoFar

	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		temp := max(curr, max(maxSoFar*curr, minSoFar*curr))
		minSoFar = min(curr, min(maxSoFar*curr, minSoFar*curr))

		maxSoFar = temp
		result = max(maxSoFar, result)
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
