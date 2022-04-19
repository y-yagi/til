package main

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		newIndex := abs(nums[i]) - 1

		if nums[newIndex] > 0 {
			nums[newIndex] *= -1
		}
	}

	result := []int{}
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			result = append(result, i)
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
