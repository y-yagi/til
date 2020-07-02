package main

// Template for 2d grid map
func main(grid [][]int) {
	m := len(grid)
	if m == 0 {
		return
	}

	n := len(grid[0])
	q := [][]int{}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			// Check value is valid or invalid
			if grid[row][col] == 0 {
				q = append(q, []int{row, col})
			}
		}
	}

	directions := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}

	for len(q) != 0 {
		point := q[0]
		q = q[1:]

		row := point[0]
		col := point[1]

		for _, direction := range directions {
			r := row + direction[0]
			c := col + direction[1]

			if r < 0 || c < 0 || r >= m || c >= n || grid[r][c] != Empty {
				continue
			}
			grid[r][c] = grid[row][col] + 1
			q = append(q, []int{r, c})
		}
	}
}
