package main

import "math"

func consecutiveNumbersSum(N int) int {
	count := 0
	// x > 0 --> N/k - (k + 1)/2 > 0
	upperLimit := int(math.Sqrt(2*float64(N)+0.25) - 0.5)
	for k := 1; k <= upperLimit; k++ {
		if (N-k*(k+1)/2)%k == 0 {
			count++
		}
	}

	return count
}
