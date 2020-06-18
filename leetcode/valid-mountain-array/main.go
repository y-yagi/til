package main

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	if A[0] > A[1] {
		return false
	}

	up := true
	for i := 1; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return false
		}

		if !up && A[i] < A[i+1] {
			return false
		}

		if up && A[i] > A[i+1] {
			up = false
		}
	}

	if up {
		return false
	}
	return true
}
