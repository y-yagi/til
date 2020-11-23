package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		d(dp)
		for x := coin; x < amount+1; x++ {
			dp[x] += dp[x-coin]
		}
	}
	d(dp)

	return dp[amount]
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
