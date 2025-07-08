package main

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, num := range nums {
		dict[target-num] = i
	}

	for i, num := range nums {
		if j, found := dict[num]; found && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}
