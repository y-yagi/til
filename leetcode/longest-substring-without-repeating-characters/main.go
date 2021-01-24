package main

func lengthOfLongestSubstring(s string) int {
	dict := map[byte]int{}
	ans := 0
	i := 0
	j := 0

	for ; j < len(s); j++ {
		if pos, found := dict[s[j]]; found {
			i = max(pos, i)
		}
		ans = max(ans, j-i+1)
		dict[s[j]] = j + 1
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
