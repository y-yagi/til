package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	queue := make([]int, 0)
	ans := 1

	for i := 0; i < len(nums); i++ {
		if len(queue) == 0 {
			queue = append(queue, nums[i])
			continue
		}

		last := queue[len(queue)-1]
		if last == nums[i] {
			continue
		} else if last == nums[i]+1 || last == nums[i]-1 {
			queue = append(queue, nums[i])
		} else {
			ans = max(ans, len(queue))
			queue = []int{nums[i]}
		}
	}

	ans = max(ans, len(queue))
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
