package main

import "fmt"

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}

func searchRange(nums []int, target int) []int {
	start := findStart(nums, target, 0, len(nums)-1)
	if start == -1 {
		return []int{-1, -1}
	}

	i := start
	for ; i < len(nums); i++ {
		if nums[i] != target {
			break
		}
	}

	return []int{start, i - 1}
}

func findStart(nums []int, target, l, r int) int {
	if l > r {
		return -1
	}

	middle := (l + r) / 2

	if nums[middle] < target {
		return findStart(nums, target, middle+1, r)
	} else if nums[middle] > target {
		return findStart(nums, target, l, middle-1)
	}

	for middle != 0 && nums[middle-1] == target {
		middle--
	}
	return middle
}
