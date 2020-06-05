package main

type pair struct {
	first  int
	second int
}

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}

	rows := len(board)
	cols := len(board[0])

	borders := make([]pair, 0)
	for i := 0; i < rows; i++ {
		borders = append(borders, pair{first: i, second: 0})
		borders = append(borders, pair{first: i, second: cols - 1})
	}

	for j := 0; j < cols; j++ {
		borders = append(borders, pair{first: 0, second: j})
		borders = append(borders, pair{first: rows - 1, second: j})
	}

	for _, p := range borders {
		dfs(board, p.first, p.second, rows, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
			if board[r][c] == 'E' {
				board[r][c] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, row, col, rows, cols int) {
	if board[row][col] != 'O' {
		return
	}

	board[row][col] = 'E'
	if col < cols-1 {
		dfs(board, row, col+1, rows, cols)
	}
	if row < rows-1 {
		dfs(board, row+1, col, rows, cols)
	}
	if col > 0 {
		dfs(board, row, col-1, rows, cols)
	}
	if row > 0 {
		dfs(board, row-1, col, rows, cols)
	}
}
