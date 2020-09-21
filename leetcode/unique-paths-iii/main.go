package main

func uniquePathsIII(grid [][]int) int {
	var startRow, startCol, nonObstacles int

	rows := len(grid)
	cols := len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			cell := grid[row][col]
			if cell >= 0 {
				nonObstacles++
			}
			if cell == 1 {
				startRow = row
				startCol = col
			}
		}
	}

	ans := 0
	backtrace(startRow, startCol, nonObstacles, rows, cols, grid, &ans)
	return ans
}

func backtrace(row, col, remain, rows, cols int, grid [][]int, ans *int) {
	if grid[row][col] == 2 && remain == 1 {
		*ans++
		return
	}

	temp := grid[row][col]
	grid[row][col] = -4
	remain--

	rowOffsets := []int{0, 0, 1, -1}
	colOffsets := []int{1, -1, 0, 0}
	for i := 0; i < 4; i++ {
		nextRow := row + rowOffsets[i]
		nextCol := col + colOffsets[i]

		if 0 > nextRow || nextRow >= rows || 0 > nextCol || nextCol >= cols {
			continue
		}

		if grid[nextRow][nextCol] < 0 {
			continue
		}

		backtrace(nextRow, nextCol, remain, rows, cols, grid, ans)
	}

	grid[row][col] = temp
}
