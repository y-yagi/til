package main

func searchMatrix(matrix [][]int, target int) bool {
	var x, y int
	for x < len(matrix) && y < len(matrix[0]) {
		if matrix[x][y] == target {
			return true
		}

		if x < len(matrix)-1 && target >= matrix[x+1][y] {
			x++
		} else {
			y++
		}
	}

	return false
}
