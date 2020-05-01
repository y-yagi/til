package main

func setZeroes(matrix [][]int) {
	changed := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		changed[i] = make([]bool, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if !changed[i][j] && matrix[i][j] == 0 {
				for k := 0; k < len(matrix[0]); k++ {
					if matrix[i][k] != 0 {
						matrix[i][k] = 0
						changed[i][k] = true
					}
				}
				for k := 0; k < len(matrix); k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = 0
						changed[k][j] = true
					}
				}
			}
		}
	}
}
