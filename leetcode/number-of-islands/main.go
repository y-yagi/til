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

func walk(i, j int, byte [][]byte) int {
	if i < 0 || j < 0 || i >= len(byte) || j >= len(byte[0]) {
		return 0
	}
	if byte[i][j] == '1' {
		byte[i][j] = '2'
		_ = walk(i-1, j, byte)
		_ = walk(i+1, j, byte)
		_ = walk(i, j-1, byte)
		_ = walk(i, j+1, byte)
		return 1
	}
	return 0
}
