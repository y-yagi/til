package main

import (
	"fmt"
	"sort"
	"strings"
)

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}

func groupAnagrams(strs []string) [][]string {
	groups := [][]string{}
	anagrams := map[string]int{}

	for _, str := range strs {
		sortedStr := sortString(str)
		if _, exist := anagrams[sortedStr]; exist {
			index := anagrams[sortedStr]
			groups[index] = append(groups[index], str)
		} else {
			groups = append(groups, []string{str})
			anagrams[sortedStr] = len(groups) - 1
		}

	}

	return groups
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
