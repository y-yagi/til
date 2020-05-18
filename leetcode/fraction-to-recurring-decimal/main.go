package main

import (
	"fmt"
	"math"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	fraction := ""

	if numerator*denominator < 0 {
		fraction += "-"
	}

	dividend := int(math.Abs(float64(numerator)))
	divisor := int(math.Abs(float64(denominator)))
	fraction += fmt.Sprintf("%d", dividend/divisor)
	remaider := dividend % divisor
	if remaider == 0 {
		return fraction
	}

	fraction += "."
	dict := map[int]int{}
	for remaider != 0 {
		if v, found := dict[remaider]; found {
			fraction = fraction[:v] + "(" + fraction[v:]
			fraction += ")"
			break
		}
		dict[remaider] = len(fraction)
		remaider *= 10
		fraction += fmt.Sprintf("%d", remaider/divisor)
		remaider %= divisor
	}

	return fraction
}
