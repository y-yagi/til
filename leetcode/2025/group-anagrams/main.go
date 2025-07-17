package main

import (
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := getSignature(s)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func getSignature(s string) string {
	count := make([]int, 26)

	for _, c := range s {
		count[c-'a']++
	}

	var sb strings.Builder
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			sb.WriteByte(byte(i + 'a'))
			sb.WriteString(strconv.Itoa(count[i]))
		}
	}

	return sb.String()
}
