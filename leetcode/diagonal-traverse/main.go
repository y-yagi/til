package main

func findDiagonalOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	N := len(matrix)
	M := len(matrix[0])
	ans := make([]int, N*M)
	k := 0

	var r, c int

	for d := 0; d < N+M-1; d++ {
		intermediate := []int{}

		// We need to figure out the "head" of this diagonal
		// The elements in the first row and the last column
		// are the respective heads.
		if d < M {
			r = 0
			c = d
		} else {
			r = d - M + 1
			c = M - 1
		}

		for r < N && c > -1 {
			intermediate = append(intermediate, matrix[r][c])
			r++
			c--
		}

		if d%2 == 0 {
			intermediate = reverse(intermediate)
		}

		for i := 0; i < len(intermediate); i++ {
			ans[k] = intermediate[i]
			k++
		}
	}

	return ans
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
