package main

import "strings"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	ans := []string{}
	for i := 0; i < len(values) && num >= 0; i++ {
		for values[i] <= num {
			num -= values[i]
			ans = append(ans, symbols[i])
		}
	}

	return strings.Join(ans, "")
}

func intToRomanFirst(num int) string {
	ans := ""
	if t := num / 1000; t > 0 {
		num = num % 1000
		ans += strings.Repeat("M", t)
	}

	ans += toRoman(num, 100, "C", "D", "M")
	num = num % 100

	ans += toRoman(num, 10, "X", "L", "C")
	num = num % 10

	ans += toRoman(num, 1, "I", "V", "X")
	return ans
}

func toRoman(num, position int, one, five, ten string) string {
	res := ""
	if t := num / position; t > 0 {
		switch t {
		case 4:
			res += one + five
		case 9:
			res += one + ten
		default:
			if t >= 5 {
				res += five
				t -= 5
			}
			res += strings.Repeat(one, t)
		}
	}
	return res
}
