package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	combinations := [][]int{}
	r := []int{}
	sort.Ints(candidates)
	find(candidates, target, &combinations, r, 0)
	return combinations
}

func find(candidates []int, target int, combinations *[][]int, r []int, i int) {
	if target < 0 {
		return
	}

	if target == 0 {
		cr := make([]int, len(r))
		copy(cr, r)
		*combinations = append(*combinations, cr)
		return
	}

	for i < len(candidates) && target-candidates[i] >= 0 {
		r = append(r, candidates[i])
		find(candidates, target-candidates[i], combinations, r, i)
		i++
		r = r[:len(r)-1]
	}

	return
}
