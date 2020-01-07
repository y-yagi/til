package main

func maxProfit(prices []int) int {
	i := 0
	valley := prices[0]
	peak := prices[0]
	profit := 0

	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]

		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]

		profit += peak - valley
	}

	return profit
}
