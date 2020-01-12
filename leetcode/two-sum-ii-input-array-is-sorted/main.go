package main

func twoSum(nums []int, target int) []int {
	start := 0
	last := len(nums) - 1

	for start < last {
		v := nums[start] + nums[last]
		if v == target {
			return []int{start + 1, last + 1}
		} else if v < target {
			start++
		} else {
			last--
		}
	}

	return []int{}
}
