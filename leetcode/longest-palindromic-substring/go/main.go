package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var start, end int
	for i := 0; i < len(s); i++ {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		l := max(l1, l2)
		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
