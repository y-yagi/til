package main

type TicTacToe struct {
	rows         []int
	cols         []int
	diagonal     int
	antiDiagonal int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	rows := make([]int, n)
	cols := make([]int, n)
	return TicTacToe{rows: rows, cols: cols}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	toAdd := 1
	if player == 2 {
		toAdd = -1
	}

	this.rows[row] += toAdd
	this.cols[col] += toAdd
	if row == col {
		this.diagonal += toAdd
	}
	if col == len(this.cols)-row-1 {
		this.antiDiagonal += toAdd
	}

	n := len(this.rows)
	if this.abs(this.rows[row]) == n || this.abs(this.cols[col]) == n || this.abs(this.diagonal) == n || this.abs(this.antiDiagonal) == n {
		return player
	}

	return 0
}

func (this *TicTacToe) abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
