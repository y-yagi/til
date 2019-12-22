package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	sum := ""
	carried := 0

	for i >= 0 || j >= 0 {
		now := carried
		if i >= 0 {
			z, _ := strconv.Atoi(string(a[i]))
			now += z
			i--
		}

		if j >= 0 {
			z, _ := strconv.Atoi(string(b[j]))
			now += z
			j--
		}

		carried = now / 2
		sum = strconv.Itoa(now%2) + sum
	}

	if carried > 0 {
		sum = strconv.Itoa(carried) + sum
	}

	return sum
}
