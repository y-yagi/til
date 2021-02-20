package main

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans = max(ans, helper(i, j, grid))
		}
	}

	return ans
}

func helper(x, y int, grid [][]int) int {
	if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[0])-1 || grid[x][y] == 0 {
		return 0
	}

	grid[x][y] = 0
	return 1 + helper(x+1, y, grid) + helper(x, y+1, grid) + helper(x-1, y, grid) + helper(x, y-1, grid)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
