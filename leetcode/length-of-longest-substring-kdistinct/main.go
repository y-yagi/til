package main

import (
	"math"
)

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	dict := map[byte]int{}
	var ans, left, right int

	for right < len(s) {
		dict[s[right]] = right
		right++

		if len(dict) == k+1 {
			var key byte
			min := math.MaxInt32
			for k, v := range dict {
				if min > v {
					min = v
					key = k
				}
			}
			delete(dict, key)
			left = min + 1
		}
		ans = max(ans, right-left)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
