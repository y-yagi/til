package main

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2+n%2; i++ {
		for j := 0; j < n/2; j++ {
			tmp := make([]int, 4)
			row := i
			col := j
			for k := 0; k < 4; k++ {
				tmp[k] = matrix[row][col]
				x := row
				row = col
				col = n - 1 - x
			}

			for k := 0; k < 4; k++ {
				matrix[row][col] = tmp[(k+3)%4]
				x := row
				row = col
				col = n - 1 - x
			}
		}
	}
}
