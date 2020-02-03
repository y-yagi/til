package main

func addDigits(num int) int {
	m := num % 9
	if m == 0 && num != 0 {
		return 9
	}
	return m
}
