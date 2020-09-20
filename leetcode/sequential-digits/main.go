package main

import (
	"strconv"
)

func sequentialDigits(low int, high int) []int {
	sample := "123456789"
	n := 10
	nums := []int{}

	lowLen := len(strconv.Itoa(low))
	highLen := len(strconv.Itoa(high))
	for length := lowLen; length < highLen+1; length++ {
		for start := 0; start < n-length; start++ {
			num, _ := strconv.Atoi(string(sample[start : start+length]))
			if num >= low && num <= high {
				nums = append(nums, num)
			}
		}
	}

	return nums
}
