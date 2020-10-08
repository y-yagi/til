package main

import "sort"

type counter struct {
	value string
	count int
}

func topKFrequent(words []string, k int) []string {
	m := map[string]int{}
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}

	counters := []counter{}
	for k, v := range m {
		counters = append(counters, counter{value: k, count: v})
	}

	sort.Slice(counters, func(i, j int) bool {
		if counters[i].count != counters[j].count {
			return counters[i].count > counters[j].count
		}

		s := []string{counters[i].value, counters[j].value}
		sort.Strings(s)
		if s[0] == counters[i].value {
			return true
		}
		return false
	})

	ans := []string{}
	for i := 0; i < k; i++ {
		ans = append(ans, counters[i].value)
	}

	return ans
}
