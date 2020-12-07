package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	emails := map[string][]int{}
	for i := range accounts {
		for j := 1; j < len(accounts[i]); j++ {
			if _, found := emails[accounts[i][j]]; !found {
				emails[accounts[i][j]] = []int{len(emails), i}
			}
		}
	}

	uf := newUnionFind(len(emails))
	for i := range accounts {
		if len(accounts[i]) < 3 {
			continue
		}
		for j := 2; j < len(accounts[i]); j++ {
			uf.union(emails[accounts[i][j-1]][0], emails[accounts[i][j]][0])
		}
	}

	groups := map[int]int{}
	merged := [][]string{}
	for email, i := range emails {
		group := uf.find(i[0])
		if _, found := groups[group]; !found {
			merged = append(merged, []string{accounts[i[1]][0], email})
			groups[group] = len(merged) - 1
		} else {
			merged[groups[group]] = append(merged[groups[group]], email)
		}
	}

	for i := range merged {
		sort.Strings(merged[i][1:])
	}

	return merged
}

type uf struct{ uf []int }

func (u uf) union(a, b int) {
	i, j := u.find(a), u.find(b)
	if i != j {
		u.uf[i] = j
	}
}

func (u uf) find(i int) int {
	if u.uf[i] == 0 {
		return i
	}
	return u.find(u.uf[i])
}

func newUnionFind(size int) uf {
	return uf{make([]int, size+1)}
}
