package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	groups := [][]string{}
	dict := map[string]int{}

	for i := 0; i < len(strs); i++ {
		b := []byte(strs[i])
		sort.Slice(b, func(i int, j int) bool { return b[i] < b[j] })

		if index, found := dict[string(b)]; found {
			groups[index] = append(groups[index], strs[i])
		} else {
			groups = append(groups, []string{strs[i]})
			dict[string(b)] = len(groups) - 1
		}
	}

	return groups
}
