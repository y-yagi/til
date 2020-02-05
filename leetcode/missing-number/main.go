package main

func missingNumber(nums []int) int {
	dict := make(map[int]bool, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		dict[nums[i]] = false
	}

	for i := 0; i < len(nums)+1; i++ {
		_, found := dict[i]
		if !found {
			return i
		}
	}
	return 0
}
