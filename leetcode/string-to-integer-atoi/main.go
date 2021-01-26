package main

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	sign := 1
	i := 0
	total := 0

	if str[0] == '+' {
		i++
	} else if str[0] == '-' {
		i++
		sign = -1
	}

	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		digit := int(str[i] - '0')

		if math.MaxInt32/10 < total || math.MaxInt32/10 == total && math.MaxInt32%10 < digit {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		total = 10*total + digit
	}

	return total * sign
}
