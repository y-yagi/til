package main

func orangesRotting(grid [][]int) int {
	rottens := [][]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				rottens = append(rottens, []int{i, j})
			}
		}
	}

	timer := 0
	for len(rottens) != 0 {
		l := len(rottens)
		for i := 0; i < l; i++ {
			rotten := rottens[0]
			rottens = rottens[1:]
			i := rotten[0]
			j := rotten[1]

			if i > 0 && grid[i-1][j] == 1 {
				grid[i-1][j] = 2
				rottens = append(rottens, []int{i - 1, j})
			}
			if j > 0 && grid[i][j-1] == 1 {
				grid[i][j-1] = 2
				rottens = append(rottens, []int{i, j - 1})
			}
			if i < len(grid)-1 && grid[i+1][j] == 1 {
				grid[i+1][j] = 2
				rottens = append(rottens, []int{i + 1, j})
			}
			if j < len(grid[0])-1 && grid[i][j+1] == 1 {
				grid[i][j+1] = 2
				rottens = append(rottens, []int{i, j + 1})
			}
		}
		timer++
	}
	if timer != 0 {
		timer--
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return timer
}
