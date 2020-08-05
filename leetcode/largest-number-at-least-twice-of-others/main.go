package main

func dominantIndex(nums []int) int {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if maxIndex != i && nums[maxIndex] < 2*nums[i] {
			return -1
		}
	}

	return maxIndex
}
