package main

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	candicates := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}

			for _, c := range candicates {
				if isValid(board, i, j, c) {
					board[i][j] = c
					if solve(board) {
						return true
					} else {
						board[i][j] = '.'
					}
				}
			}
			return false
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' && board[i][col] == c {
			return false
		}
		if board[row][i] != '.' && board[row][i] == c {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] != '.' && board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}
	return true
}
