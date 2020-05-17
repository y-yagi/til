package main

import "math"

const HALF_INT_MIN = -1073741824

func divide(dividend int, divisor int) int {
	if dividend == math.MaxInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	/* We need to convert both numbers to negatives
	 * for the reasons explained above.
	 * Also, we count the number of negatives signs. */
	negatives := 2
	if dividend > 0 {
		negatives--
		dividend = -dividend
	}
	if divisor > 0 {
		negatives--
		divisor = -divisor
	}

	/* Count how many times the divisor has to be added
	 * to get the dividend. This is the quotient. */
	quotient := 0
	for divisor >= dividend {
		powerOfTwo := -1
		value := divisor

		for value >= HALF_INT_MIN && value+value >= dividend {
			value += value
			powerOfTwo += powerOfTwo
		}
		quotient += powerOfTwo
		dividend -= value
	}

	/* If there was originally one negative sign, then
	 * the quotient remains negative. Otherwise, switch
	 * it to positive. */
	if negatives != 1 {
		return -quotient
	}

	return quotient
}
