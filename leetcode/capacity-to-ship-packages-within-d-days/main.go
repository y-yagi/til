package main

import "fmt"

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}

func shipWithinDays(weights []int, D int) int {
	var low, high int
	for _, w := range weights {
		low = max(low, w)
		high += w
	}

	for low < high {
		mid := (low + high) / 2
		total := 0
		days := 1

		for _, w := range weights {
			if total+w > mid {
				days += 1
				total = w
			} else {
				total += w
			}
		}

		if days <= D {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
