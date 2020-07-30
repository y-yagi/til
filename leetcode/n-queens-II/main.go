package main

func totalNQueens(n int) int {
	rows := make([]int, n)
	hills := make([]int, 4*n-1)
	dales := make([]int, 2*n-1)

	return backtrace(0, 0, n, rows, hills, dales)
}

func backtrace(row, count, n int, rows, hills, dales []int) int {
	for col := 0; col < n; col++ {
		if isNotUnderAttack(row, col, n, rows, hills, dales) {
			rows[col] = 1
			hills[row-col+2*n] = 1
			dales[row+col] = 1

			if row+1 == n {
				// if n queens are already placed
				count++
			} else {
				// if not proceed to place the rest
				count = backtrace(row+1, count, n, rows, hills, dales)
			}

			rows[col] = 0
			hills[row-col+2*n] = 0
			dales[row+col] = 0
		}
	}
	return count
}

func isNotUnderAttack(row, col, n int, rows, hills, dales []int) bool {
	res := rows[col] + hills[row-col+2*n] + dales[row+col]

	if res == 0 {
		return true
	}
	return false
}
