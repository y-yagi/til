package main

func stoneGame(piles []int) bool {
	return true
}

func stoneGame_dp(piles []int) bool {
	n := len(piles)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for size := 1; size <= n; size++ {
		for i := 0; i+size <= n; i++ {
			j := i + size - 1
			parity := (j + i + n) % 2
			if parity == 1 {
				dp[i+1][j+1] = max(piles[i]+dp[i+2][j+1], piles[j]+dp[i+1][j])
			} else {
				dp[i+1][j+1] = min(-piles[i]+dp[i+2][j+1], -piles[j]+dp[i+1][j])
			}

		}
	}

	return dp[1][n] > 0
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
