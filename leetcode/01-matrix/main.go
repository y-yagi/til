package main

import "math"

func updateMatrix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = math.MaxInt16
			}
		}
	}
	maxI, maxJ := len(matrix)-1, len(matrix[0])-1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				matrix[i][j] = minPoint(matrix[i][j], i-1, j, maxI, maxJ, 1, matrix)
				matrix[i][j] = minPoint(matrix[i][j], i, j-1, maxI, maxJ, 1, matrix)
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][j] != 0 {
				matrix[i][j] = minPoint(matrix[i][j], i, j+1, maxI, maxJ, 1, matrix)
				matrix[i][j] = minPoint(matrix[i][j], i+1, j, maxI, maxJ, 1, matrix)
			}
		}
	}
	return matrix
}

func minPoint(v, i, j, maxI, maxJ, more int, matrix [][]int) int {
	if i < 0 || j < 0 || i > maxI || j > maxJ {
		return v
	}
	if v < matrix[i][j]+more {
		return v
	}
	return matrix[i][j] + more
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
