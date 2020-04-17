package main

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[int]bool, 9)
	columns := make([]map[int]bool, 9)
	boxes := make([]map[int]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]bool, 9)
		columns[i] = make(map[int]bool, 9)
		boxes[i] = make(map[int]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			n := int(board[i][j])
			boxIndex := (i/3)*3 + j/3

			if _, found := rows[i][n]; found {
				return false
			}
			if _, found := columns[j][n]; found {
				return false
			}
			if _, found := boxes[boxIndex][n]; found {
				return false
			}

			rows[i][n] = true
			columns[j][n] = true
			boxes[boxIndex][n] = true
		}
	}

	return true
}
