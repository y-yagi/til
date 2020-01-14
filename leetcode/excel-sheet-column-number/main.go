package main

import (
	"math"
	"strings"
)

func titleToNumber(s string) int {
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	l := len(a)
	answer := 0
	piler := 0

	for i := len(s) - 1; i >= 0; i-- {
		c := math.Pow(float64(l), float64(piler))
		answer += int(c) * (strings.Index(a, string(s[i])) + 1)
		piler++
	}

	return answer
}
