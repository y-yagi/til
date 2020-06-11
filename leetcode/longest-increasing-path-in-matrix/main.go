package main

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	ans := 0
	cache := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		cache = append(cache, make([]int, len(matrix[0])))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ans = max(ans, dfs(i, j, matrix, cache))
		}
	}
	return ans
}

func dfs(row, col int, matrix, cache [][]int) int {
	if cache[row][col] != 0 {
		return cache[row][col]
	}

	for _, value := range [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}} {
		x := row + value[0]
		y := col + value[1]

		if 0 <= x && x < len(matrix) && 0 <= y && y < len(matrix[0]) && matrix[x][y] > matrix[row][col] {
			cache[row][col] = max(cache[row][col], dfs(x, y, matrix, cache))
		}
	}

	cache[row][col]++
	return cache[row][col]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
