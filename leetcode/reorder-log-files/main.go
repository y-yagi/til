package main

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	letterlogs := []string{}
	digitogs := []string{}

	for _, log := range logs {
		splited := strings.Split(log, " ")
		if unicode.IsDigit([]rune(splited[1])[0]) {
			digitogs = append(digitogs, log)
		} else {
			letterlogs = append(letterlogs, log)
		}
	}

	sort.Slice(letterlogs, func(i, j int) bool {
		s1 := strings.Split(letterlogs[i], " ")
		s2 := strings.Split(letterlogs[j], " ")
		for i := 1; i < len(s1); i++ {
			comapred := strings.Compare(s1[i], s2[i])
			if comapred == 1 {
				return false
			} else if comapred == -1 {
				return true
			}
		}
		return true
	})

	return append(letterlogs, digitogs...)
}
