package main

func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		switch c {
		case "I":
			if len(s) > i+1 && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
				c = string(s[i+1])
				if c == "V" {
					sum += 4
				} else if c == "X" {
					sum += 9
				}
				i++
			} else {
				sum += 1
			}
		case "V":
			sum += 5
		case "X":
			if len(s) > i+1 && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
				c = string(s[i+1])
				if c == "L" {
					sum += 40
				} else if c == "C" {
					sum += 90
				}
				i++
			} else {
				sum += 10
			}
		case "L":
			sum += 50
		case "C":
			if len(s) > i+1 && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
				c = string(s[i+1])
				if c == "D" {
					sum += 400
				} else if c == "M" {
					sum += 900
				}
				i++
			} else {
				sum += 100
			}
		case "D":
			sum += 500
		case "M":
			sum += 1000
		}
	}

	return sum
}
