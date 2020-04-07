package main

import "fmt"

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}

func maxAreaOfIsland(grid [][]int) int {
	seen := [][]bool{}
	max := 0
	for i := 0; i < len(grid); i++ {
		seen = append(seen, make([]bool, len(grid[0])))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			v := area(i, j, grid, seen)
			if v > max {
				max = v
			}
		}
	}

	return max
}

func area(i, j int, grid [][]int, seen [][]bool) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || seen[i][j] || grid[i][j] == 0 {
		return 0
	}

	seen[i][j] = true
	return 1 + area(i+1, j, grid, seen) + area(i-1, j, grid, seen) + area(i, j+1, grid, seen) + area(i, j-1, grid, seen)
}
