package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	versions1 := strings.Split(version1, ".")
	versions2 := strings.Split(version2, ".")
	l1 := len(versions1)
	l2 := len(versions2)

	for i := 0; i < max(l1, l2); i++ {
		var i1, i2 int
		if i < l1 {
			i1, _ = strconv.Atoi(versions1[i])
		}
		if i < l2 {
			i2, _ = strconv.Atoi(versions2[i])
		}
		if i1 > i2 {
			return 1
		} else if i1 < i2 {
			return -1
		}
	}

	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
