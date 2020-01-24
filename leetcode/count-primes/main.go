package main

import "math"

func countPrimes(n int) int {
	count := 0
	f := make([]bool, n)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if f[i] == false {
			for j := i * i; j < n; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if f[i] == false {
			count++
		}
	}
	return count
}
