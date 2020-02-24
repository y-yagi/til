package main

import (
	"strings"
)

func countSegments(s string) int {
	a := strings.Split(s, " ")
	l := len(a)

	for _, v := range a {
		if v == "" {
			l--
		}
	}

	return l
}
