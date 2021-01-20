package main

import (
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		return gcdOfStrings(str2, str1)
	} else if !strings.HasPrefix(str1, str2) { // shorter string is not common prefix.
		return ""
	} else if len(str2) == 0 {
		return str1
	}

	return gcdOfStrings(str1[len(str2):], str2)
}
