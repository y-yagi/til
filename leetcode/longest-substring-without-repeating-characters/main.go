package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", lengthOfLongestSubstring("dvdf"))
	fmt.Printf("%d\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("%d\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("%d\n", lengthOfLongestSubstring("pwwkew"))
	fmt.Printf("%d\n", lengthOfLongestSubstring("ohvhjdml"))
}

func lengthOfLongestSubstring(s string) int {
	dict := [128]bool{}
	length, max := 0, 0
	for i, j := 0, 0; i < len(s); i++ {
		index := s[i]
		if dict[index] {
			for ; dict[index]; j++ {
				length--
				dict[s[j]] = false
			}
		}

		dict[index] = true
		length++
		if length > max {
			max = length
		}
	}
	return max
}
