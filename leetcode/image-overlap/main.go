package main

func largestOverlap(A [][]int, B [][]int) int {
	n := len(A)
	largest := 0

	for h := 1; h <= n; h++ {
		for w := 1; w <= n; w++ {
			largest = max(largest, overlap(A, B, h, w))
			largest = max(largest, overlap(B, A, h, w))
		}
	}

	return largest
}

func overlap(A, B [][]int, h, w int) int {
	n := len(A)
	count := 0

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if B[i][j] == 1 && B[i][j] == A[n-h+i][n-w+j] {
				count++
			}
		}
	}

	return count
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
