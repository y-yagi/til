package main

func gameOfLife(board [][]int) {
	// Neighbors array to find 8 neighboring cells for a given cell
	neighbors := []int{0, 1, -1}
	rows := len(board)
	cols := len(board[0])

	copyBoard := make([][]int, 0)
	for i := 0; i < rows; i++ {
		copyBoard = append(copyBoard, make([]int, cols))
		for j := 0; j < cols; j++ {
			copyBoard[i][j] = board[i][j]
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveNeighbors := 0

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if !(neighbors[i] == 0 && neighbors[j] == 0) {
						r := row + neighbors[i]
						c := col + neighbors[j]

						if (r < rows && r >= 0) && (c < cols && c >= 0) && (copyBoard[r][c] == 1) {
							liveNeighbors++
						}
					}
				}
			}

			if (copyBoard[row][col] == 1) && (liveNeighbors < 2 || liveNeighbors > 3) {
				board[row][col] = 0
			}
			if copyBoard[row][col] == 0 && liveNeighbors == 3 {
				board[row][col] = 1
			}
		}
	}
}
