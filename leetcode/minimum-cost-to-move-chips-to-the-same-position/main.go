package main

func minCostToMoveChips(position []int) int {
	even := 0
	odd := 0
	for _, i := range position {
		if i%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	return min(even, odd)
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
