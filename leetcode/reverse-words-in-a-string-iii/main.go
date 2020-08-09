package main

import (
	"strings"
)

func reverseWords(s string) string {
	a := strings.Split(s, " ")
	for i := 0; i < len(a); i++ {
		a[i] = reverse(a[i])
	}
	return strings.Join(a, " ")
}

func reverse(s string) string {
	a := []byte(s)
	l := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		t := a[i]
		a[i] = a[l-i]
		a[l-i] = t
	}
	return string(a)
}
