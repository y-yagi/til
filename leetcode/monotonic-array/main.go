package main

func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}

	compare := 0
	i := 1
	for ; i < len(A); i++ {
		if A[i-1] < A[i] {
			compare = 1
			break
		} else if A[i-1] > A[i] {
			compare = -1
			break
		}
	}

	for ; i < len(A); i++ {
		if compare == 1 && A[i-1] > A[i] {
			return false
		} else if compare == -1 && A[i-1] < A[i] {
			return false
		}
	}

	return true
}
