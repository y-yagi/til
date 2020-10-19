package main

func minDominoRotations(A []int, B []int) int {
	if numA, ok := rotateNum(A[0], A, B); ok {
		return numA
	} else if numB, ok := rotateNum(B[0], A, B); ok {
		return numB
	}
	return -1
}

func rotateNum(n int, A []int, B []int) (int, bool) {
	a, b := 0, 0
	for i := 0; i < len(A); i++ {
		if A[i] != n && B[i] != n {
			return -1, false
		} else if B[i] == n && A[i] != n {
			a++
		} else if A[i] == n && B[i] != n {
			b++
		}
	}
	return min(a, b), true
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
