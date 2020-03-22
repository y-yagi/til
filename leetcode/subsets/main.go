package main

import (
	"math"
	"strconv"
)

func subsets(nums []int) [][]int {
	answer := [][]int{}
	n := float64(len(nums))
	for i := int(math.Pow(2, n)); i < int(math.Pow(2, n+1)); i++ {
		bitmask := strconv.FormatInt(int64(i), 2)
		bitmask = bitmask[1:]

		curr := []int{}
		for j := 0; j < int(n); j++ {
			if bitmask[j] == '1' {
				curr = append(curr, nums[j])
			}
		}
		answer = append(answer, curr)
	}

	return answer
}

func subsets_Recursion(nums []int) [][]int {
	answer := [][]int{}
	answer = append(answer, []int{})

	for _, n := range nums {
		newSubsets := [][]int{}
		for _, curr := range answer {
			a := make([]int, len(curr))
			copy(a, curr)
			a = append(a, n)
			newSubsets = append(newSubsets, a)
		}

		for _, curr := range newSubsets {
			answer = append(answer, curr)
		}
	}

	return answer
}
