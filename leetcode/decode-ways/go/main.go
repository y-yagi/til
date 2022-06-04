package main

import "strconv"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1

	// Ways to decode a string of size 1 is 1. Unless the string is '0'.
	// '0' doesn't have a single digit decode.
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i < len(dp); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		twoDigit, _ := strconv.Atoi(s[i-2 : i])
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
