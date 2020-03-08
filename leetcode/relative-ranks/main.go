package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	answer := []string{}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	dict := map[int]string{0: "Gold Medal", 1: "Silver Medal", 2: "Bronze Medal"}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(sorted); j++ {
			if nums[i] == sorted[j] {
				if s, f := dict[j]; f {
					answer = append(answer, s)
				} else {
					answer = append(answer, strconv.Itoa(j+1))
				}
				break
			}
		}
	}
	return answer
}
