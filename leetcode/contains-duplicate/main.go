package main

func containsDuplicate(nums []int) bool {
	dict := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := dict[nums[i]]; ok {
			return true
		}
		dict[nums[i]] = true

	}
	return false
}
