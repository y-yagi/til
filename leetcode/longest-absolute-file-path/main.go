package main

import (
	"fmt"
	"strings"
)

func lengthLongestPath(input string) int {
	var max int
	ss := strings.Split(input, "\n")
	segLen := make(map[int]int)

	for _, str := range ss {
		curIndent := getIndent(str)
		name := str[curIndent:]
		if strings.Contains(name, ".") {
			curLen := len(name)
			for i := curIndent - 1; i >= 0; i-- {
				curLen += segLen[i] + 1
			}
			if curLen > max {
				max = curLen
			}
		} else {
			segLen[curIndent] = len(name)
		}
	}
	return max
}

func getIndent(str string) int {
	var val int
	for i := 0; i < len(str); i++ {
		if str[i] != '\t' {
			break
		}
		val++
	}
	return val
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
