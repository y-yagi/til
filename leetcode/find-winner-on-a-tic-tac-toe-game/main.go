package main

func tictactoe(moves [][]int) string {
	board := [3][3]string{}
	for i := 0; i < len(moves); i++ {
		row := moves[i][0]
		col := moves[i][1]
		if len(moves)%2 == 0 {
			board[row][col] = "X"
			if didWin(board, row, col, "X") {
				return "A"
			}
		} else {
			if didWin(board, row, col, "O") {
				return "B"
			}
		}
	}

	if len(board) != 3*3 {
		return "Pending"
	}

	return "Draw"
}

func didWin(board [3][3]string, row, col int, player string) bool {
	didPlayerWin := true
	for i := 0; i < 3; i++ {
		if board[row][i] != player {
			didPlayerWin = false
		}
	}
	if didPlayerWin {
		return true
	}

	// check the current col
	didPlayerWin = true
	for i := 0; i < 3; i++ {
		if board[i][col] != player {
			didPlayerWin = false
		}
	}
	if didPlayerWin {
		return true
	}

	//check the diagonal
	if row == col {
		didPlayerWin = true
		for i := 0; i < 3; i++ {
			if board[i][i] != player {
				didPlayerWin = false
			}
		}
		if didPlayerWin {
			return true
		}
	}

	//check reverse diagonal
	if col == 3-row-1 {
		didPlayerWin = true
		for i := 0; i < 3; i++ {
			if board[i][3-i-1] != player {
				didPlayerWin = false
			}
		}
		if didPlayerWin {
			return true
		}
	}

	//player did not win
	return false
}
