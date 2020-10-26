package main

import "math"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	A := [101][101]float64{}
	A[0][0] = float64(poured)

	for r := 0; r <= query_row; r++ {
		for c := 0; c <= r; c++ {
			q := (A[r][c] - 1.0) / 2.0
			if q > 0 {
				A[r+1][c] += q
				A[r+1][c+1] += q
			}
		}
	}

	return math.Min(1.0, A[query_row][query_glass])
}
