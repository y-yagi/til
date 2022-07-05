package main

import "strings"

func minRemoveToMakeValid(s string) string {
	stack := []int{}
	b := []byte(s)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				b[i] = '*'
			}
		}
	}

	for len(stack) > 0 {
		b[stack[len(stack)-1]] = '*'
		stack = stack[:len(stack)-1]
	}

	return strings.ReplaceAll(string(b), "*", "")
}
