package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	ans := [][]int{}
	sort.Ints(arr)
	min := math.MaxInt32
	for i := 1; i < len(arr); i++ {
		t := arr[i] - arr[i-1]
		if t < min {
			ans = [][]int{}
			ans = append(ans, []int{arr[i-1], arr[i]})
			min = t
		} else if t == min {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}

	return ans
}
