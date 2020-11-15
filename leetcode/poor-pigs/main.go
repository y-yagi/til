package main

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// Ref: https://en.wikipedia.org/wiki/Qubit
	states := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(states))))
}
