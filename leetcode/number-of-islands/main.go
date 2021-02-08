package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	islands := 0
	for i := range grid {
		for j := range grid[i] {
			islands += walk(i, j, grid)
		}
	}

	return islands
}

func walk(i, j int, grid [][]byte) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '2' {
		return 0
	}
	if grid[i][j] == '1' {
		grid[i][j] = '2'
		walk(i-1, j, grid)
		walk(i+1, j, grid)
		walk(i, j-1, grid)
		walk(i, j+1, grid)
		return 1
	}
	return 0
}
