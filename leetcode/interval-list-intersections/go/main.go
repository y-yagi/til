package main

func intervalIntersection(A [][]int, B [][]int) [][]int {
	ans := [][]int{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		lo := max(A[i][0], B[j][0])
		hi := min(A[i][1], B[j][1])
		if lo <= hi {
			ans = append(ans, []int{lo, hi})
		}

		// Remove the interval with the smallest endpoint
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
