package main

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	radius := 0
	sort.Ints(houses)
	sort.Ints(heaters)

	for i, j := 0, 0; i < len(houses); i++ {
		for ; j < len(heaters)-1 && heaters[j]+heaters[j+1] <= 2*houses[i]; j++ {
		}
		radius = int(math.Max(float64(radius), math.Abs(float64(heaters[j]-houses[i]))))
	}
	return radius
}
