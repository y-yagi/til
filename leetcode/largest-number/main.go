package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(strs, func(i, j int) bool {
		order1, _ := strconv.Atoi(strs[i] + strs[j])
		order2, _ := strconv.Atoi(strs[j] + strs[i])
		if order1 > order2 {
			return true
		}
		return false
	})

	if strs[0] == "0" {
		return "0"
	}

	ans := ""
	for _, str := range strs {
		ans += str
	}

	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
