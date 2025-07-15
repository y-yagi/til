package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	cur := []int{}
	sort.Ints(candidates)
	backtrack(candidates, target, cur, &ans, 0)
	return ans
}

func backtrack(candidates []int, target int, cur []int, ans *[][]int, i int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, cur...))
		return
	}

	for i < len(candidates) && target-candidates[i] >= 0 {
		cur = append(cur, candidates[i])
		backtrack(candidates, target-candidates[i], cur, ans, i)
		i++
		cur = cur[:len(cur)-1]
	}
}
