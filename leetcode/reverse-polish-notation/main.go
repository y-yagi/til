package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	var x, y int

	for _, token := range tokens {
		switch token {
		case "+":
			x, y, stack = fetchValues(stack)
			stack = append(stack, y+x)
		case "-":
			x, y, stack = fetchValues(stack)
			stack = append(stack, y-x)
		case "*":
			x, y, stack = fetchValues(stack)
			stack = append(stack, y*x)
		case "/":
			x, y, stack = fetchValues(stack)
			stack = append(stack, y/x)
		default:
			value, _ := strconv.Atoi(token)
			stack = append(stack, value)
		}

	}

	return stack[0]
}

func fetchValues(stack []int) (int, int, []int) {
	var x, y int

	x, stack = stack[len(stack)-1], stack[:len(stack)-1]
	y, stack = stack[len(stack)-1], stack[:len(stack)-1]

	return x, y, stack
}
