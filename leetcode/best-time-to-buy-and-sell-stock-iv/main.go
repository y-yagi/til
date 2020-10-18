package main

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)

	if n <= 0 || k <= 0 {
		return 0
	}

	if 2*k > n {
		res := 0
		for i := 1; i < n; i++ {
			res += max(0, prices[i]-prices[i-1])
		}
		return res
	}

	dp := make([][]int, k+1)
	for t := 0; t < len(dp); t++ {
		dp[t] = make([]int, len(prices))
	}

	for t := 1; t < k+1; t++ {
		maxThusFar := math.MinInt32
		for d := 1; d < n; d++ {
			maxThusFar = max(maxThusFar, dp[t-1][d-1]-prices[d-1])
			dp[t][d] = max(dp[t][d-1], maxThusFar+prices[d])
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
