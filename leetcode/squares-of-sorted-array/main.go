package main

func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	l := len(A)
	j := 0
	for j < l && A[j] < 0 {
		j++
	}
	i := j - 1
	t := 0

	for i >= 0 && j < l {
		if A[i]*A[i] < A[j]*A[j] {
			ans[t] = A[i] * A[i]
			i--
		} else {
			ans[t] = A[j] * A[j]
			j++
		}
		t++
	}

	for i >= 0 {
		ans[t] = A[i] * A[i]
		i--
		t++
	}

	for j < l {
		ans[t] = A[j] * A[j]
		j++
		t++
	}

	return ans
}
