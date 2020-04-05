package main

import (
	"sort"
)

type pair struct {
	value []int
	sum   int
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pairs := []pair{}
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			pair := pair{value: []int{n1, n2}, sum: n1 + n2}
			pairs = append(pairs, pair)
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].sum < pairs[j].sum
	})

	answer := [][]int{}
	l := len(pairs)
	if l > k {
		l = k
	}
	for i := 0; i < l; i++ {
		answer = append(answer, pairs[i].value)
	}
	return answer
}
