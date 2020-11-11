package main

import (
	"sort"
	"strings"
)

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	order := make([]int, len(username))
	for i := range order {
		order[i] = i
	}

	sort.Slice(order, func(i, j int) bool {
		if username[order[i]] == username[order[j]] {
			return timestamp[order[i]] < timestamp[order[j]]
		} else {
			return username[order[i]] < username[order[j]]
		}
	})

	threeSequence := make(map[string]int)
	threeSequenceUser := make(map[string]string)
	maxThree := ""
	max := 0

	for i := 0; i < len(username); i++ {
		currentUser := username[order[i]]
		for j := i + 1; j < len(username) && currentUser == username[order[j]]; j++ {
			for k := j + 1; k < len(username) && currentUser == username[order[k]]; k++ {
				seq := website[order[i]] + ":" + website[order[j]] + ":" + website[order[k]]
				if threeSequenceUser[seq] == currentUser {
					continue
				} else {
					threeSequenceUser[seq] = currentUser
				}

				threeSequence[seq] = threeSequence[seq] + 1
				if threeSequence[seq] > max {
					max = threeSequence[seq]
					maxThree = seq
				} else if threeSequence[seq] == max {
					if seq < maxThree {
						max = threeSequence[seq]
						maxThree = seq
					}
				}
			}

		}
	}

	return strings.Split(maxThree, ":")
}
