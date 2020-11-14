package main

import (
	"fmt"
	"sort"
)

func reorganizeString(S string) string {
	counts := make([]int, 26)

	for i := 0; i < len(S); i++ {
		counts[S[i]-'a'] += 100
	}

	for i := 0; i < 26; i++ {
		counts[i] += i
	}
	sort.Ints(counts)

	ans := make([]byte, len(S))
	t := 1
	for _, code := range counts {
		ct := code / 100
		ch := byte('a' + (code % 100))
		if ct > (len(S)+1)/2 {
			return ""
		}

		for i := 0; i < ct; i++ {
			if t >= len(S) {
				t = 0
			}
			ans[t] = ch
			t += 2
		}
	}

	return string(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
