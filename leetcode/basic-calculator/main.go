package main

import "unicode"

func calculate(s string) int {
	stack := []int{}
	operand := 0
	result := 0
	sign := 1
	n := 1

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if unicode.IsDigit(rune(ch)) {
			operand = 10*operand + int(ch-'0')
		} else if ch == '+' {
			result += sign * operand
			sign = 1
			operand = 0
		} else if ch == '-' {
			result += sign * operand
			sign = -1
			operand = 0
		} else if ch == '(' {
			stack = append(stack, result)
			stack = append(stack, sign)

			sign = 1
			result = 0
		} else if ch == ')' {
			result += sign * operand
			n, stack = stack[len(stack)-1], stack[:len(stack)-1]
			result *= n

			n, stack = stack[len(stack)-1], stack[:len(stack)-1]
			result += n

			operand = 0
		}

	}

	return result + (sign * operand)
}
