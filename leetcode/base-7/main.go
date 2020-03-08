package main

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	seven := ""
	minus := false

	if num < 0 {
		minus = true
		num = num * -1
	}

	for num > 0 {
		seven = strconv.Itoa(num%7) + seven
		num /= 7
	}
	if minus {
		seven = "-" + seven
	}

	return seven
}
