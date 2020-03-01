package main

import "math/bits"

func hammingDistance(x int, y int) int {
	z := x ^ y
	return int(bits.OnesCount(uint(z)))
}
