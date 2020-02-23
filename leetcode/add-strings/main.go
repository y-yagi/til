package main

func addStrings(num1 string, num2 string) string {
	answer := ""
	up := false

	l1 := len(num1) - 1
	l2 := len(num2) - 1

	for l1 > -1 || l2 > -1 {
		var sum byte
		if l1 > -1 {
			sum += num1[l1] - '0'
		}
		if l2 > -1 {
			sum += num2[l2] - '0'
		}

		if up {
			sum += '1' - '0'
		}

		if sum > '9'-'0' {
			up = true
			sum -= ':' - '0'
		} else {
			up = false
		}
		answer = string(sum+'0') + answer

		l1--
		l2--
	}

	if up {
		answer = "1" + answer
	}

	return answer
}
