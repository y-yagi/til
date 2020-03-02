package main

func islandPerimeter(grid [][]int) int {
	island := 0
	neighbours := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				island++
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					neighbours++
				}
				if j < len(grid[i])-1 && grid[i][j+1] == 1 {
					neighbours++
				}
			}
		}
	}

	return island*4 - neighbours*2
}
