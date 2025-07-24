package main

import "math"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	min := math.MaxInt
	ans := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > ans {
			ans = prices[i] - min
		}
	}

	return ans
}
