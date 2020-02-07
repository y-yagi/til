package main

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	dict := map[byte]string{}
	rdict := map[string]byte{}
	q := strings.Split(str, " ")

	if len(pattern) != len(q) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if v, found := dict[pattern[i]]; found && v != q[i] {
			return false
		}
		if rv, found := rdict[q[i]]; found && rv != pattern[i] {
			return false
		}

		dict[pattern[i]] = q[i]
		rdict[q[i]] = pattern[i]
	}

	return true
}
