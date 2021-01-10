package main

func minCostClimbingStairs(cost []int) int {
	f1 := 0
	f2 := 0
	for i := len(cost) - 1; i >= 0; i-- {
		f0 := cost[i] + min(f1, f2)
		f2 = f1
		f1 = f0
	}
	return min(f1, f2)
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
