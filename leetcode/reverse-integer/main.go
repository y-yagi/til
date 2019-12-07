package main

import (
	"strconv"
)

func reverse(x int) int {
	s := strconv.Itoa(x)
	anss := ""
	if string(s[0]) == "-" {
		anss += "-"
		s = s[1:]
	}

	for i := len(s) - 1; i > -1; i-- {
		anss += string(s[i])
	}

	ans, err := strconv.ParseInt(anss, 10, 32)
	if err != nil {
		return 0
	}

	return int(ans)
}
