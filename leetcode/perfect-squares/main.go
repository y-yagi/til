package main

import (
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	maxSquareIndex := int(math.Sqrt(float64(n))) + 1
	squareNums := make([]int, maxSquareIndex)
	for i := 0; i < maxSquareIndex; i++ {
		squareNums[i] = i * i
	}

	for i := 1; i <= n; i++ {
		for s := 1; s < maxSquareIndex; s++ {
			if i < squareNums[s] {
				break
			}
			dp[i] = min(dp[i], dp[i-squareNums[s]]+1)
		}
	}

	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
