package main

import "math"

func maxProfit(prices []int) int {
	t1Cost := math.MaxInt32
	t2Cost := math.MaxInt32

	t1Profit := 0
	t2Profit := 0

	for _, price := range prices {
		t1Cost = min(t1Cost, price)
		t1Profit = max(t1Profit, price-t1Cost)
		t2Cost = min(t2Cost, price-t1Profit)
		t2Profit = max(t2Profit, price-t2Cost)
	}

	return t2Profit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
