package main

import (
	"sort"
)

type kv struct {
	id  int
	sum int
}

func highFive(items [][]int) [][]int {
	scores := map[int][]int{}
	for _, item := range items {
		scores[item[0]] = append(scores[item[0]], item[1])
	}

	var ss []kv
	for k, v := range scores {
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
		ss = append(ss, kv{k, sum(v[0:5]) / 5})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].id < ss[j].id
	})

	ans := [][]int{}
	for _, s := range ss {
		ans = append(ans, []int{s.id, s.sum})
	}

	return ans
}

func sum(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}

	return total
}
