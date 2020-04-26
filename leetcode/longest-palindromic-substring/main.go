package main

import "fmt"

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	paindromes := []string{string(s[0])}

	for i := 0; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if paindrome(s[j : i+1]) {
				paindromes = append(paindromes, s[j:i+1])
			}
		}
	}

	answer := ""
	for _, paindrome := range paindromes {
		answer = max(answer, paindrome)
	}

	return answer
}

func max(x, y string) string {
	if len(x) > len(y) {
		return x
	}

	return y
}

func paindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
