package main

func twoSum(nums []int, target int) []int {
	dict := map[int]int{}

	for i := 0; i < len(nums); i++ {
		dict[target-nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if v, found := dict[nums[i]]; found && v != i {
			return []int{i, v}
		}
	}

	return []int{}
}
