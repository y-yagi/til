package main

func exist(board [][]byte, word string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if ret, _ := backtrace(row, col, word, 0, board); ret == true {
				return true
			}
		}
	}
	return false
}

func backtrace(row, col int, word string, index int, board [][]byte) (bool, [][]byte) {
	if index >= len(word) {
		return true, board
	}

	if row < 0 || row == len(board) || col < 0 || col == len(board[0]) || board[row][col] != word[index] {
		return false, board
	}

	ret := false
	board[row][col] = '#'
	rowOffsets := []int{0, 1, 0, -1}
	colOffsets := []int{1, 0, -1, 0}

	for d := 0; d < 4; d++ {
		ret, board = backtrace(row+rowOffsets[d], col+colOffsets[d], word, index+1, board)
		if ret {
			break
		}
	}

	board[row][col] = word[index]
	return ret, board
}
