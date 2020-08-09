package main

import (
	"strings"
)

func reverseWords(s string) string {
	a := strings.Split(s, " ")
	for i := 0; i < len(a); i++ {
		if strings.Trim(a[i], "") == "" {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}

	l := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		t := a[i]
		a[i] = a[l-i]
		a[l-i] = t
	}
	return strings.Join(a, " ")
}
