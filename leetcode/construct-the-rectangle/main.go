package main

import "math"

func constructRectangle(area int) []int {
	l := area
	w := 1

	newW := 2

	maxW := int(math.Sqrt(float64(area)))

	for newW <= maxW {
		if area%newW == 0 {
			w = newW
			l = area / w
		}
		newW++
	}

	return []int{l, w}
}
