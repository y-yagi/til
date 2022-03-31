package main

import "sort"

type entry struct {
	key   int
	count int
}

func topKFrequent(nums []int, k int) []int {
	dict := map[int]int{}

	for _, num := range nums {
		dict[num]++
	}

	var s []entry
	for key, val := range dict {
		s = append(s, entry{key: key, count: val})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].count > s[j].count
	})

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, s[i].key)
	}

	return res
}
