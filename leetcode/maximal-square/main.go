package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := 0
	if rows > 0 {
		cols = len(matrix[0])
	}

	dp := [][]int{}
	for i := 0; i < rows+1; i++ {
		dp = append(dp, make([]int, cols+1))
	}

	maxsqlen := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
				maxsqlen = max(maxsqlen, dp[i][j])
			}
		}
	}

	return maxsqlen * maxsqlen
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
