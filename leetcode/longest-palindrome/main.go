package main

func longestPalindrome(s string) int {
	dict := make(map[byte]int, len(s))
	center := false
	answer := 0

	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	for _, v := range dict {
		if v%2 == 0 {
			answer += v
			continue
		}

		if v > 2 {
			answer += v - 1
			v = 1
		}

		if v == 1 && !center {
			answer += v
			center = true
			continue
		}
	}

	return answer
}
