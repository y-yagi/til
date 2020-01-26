package main

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]bool, k)
	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if set[nums[i]] {
			return true
		}
		set[nums[i]] = true
	}
	return false
}
