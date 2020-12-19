package main

import "math"

func minKnightMoves(x int, y int) int {
	x = abs(x)
	y = abs(y)
	if x < y {
		temp := x
		x = y
		y = temp
	}

	if x == 1 && y == 0 {
		return 3
	}

	if x == 2 && y == 2 {
		return 4
	}

	delta := x - y
	if y > delta {
		return (delta - 2*int(math.Floor(float64(delta-y)/3)))
	} else {
		return (delta - 2*int(math.Floor(float64((delta-y)/4))))
	}
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}
