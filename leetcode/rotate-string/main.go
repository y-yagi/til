package main

func rotateString(A string, B string) bool {
	N := len(A)
	if N != len(B) {
		return false
	}
	if N == 0 {
		return true
	}

	shifts := make([]int, N+1)
	for i := 0; i < len(shifts); i++ {
		shifts[i] = 1
	}

	left := -1

	for right := 0; right < N; right++ {
		for left >= 0 && B[left] != B[right] {
			left -= shifts[left]
		}
		shifts[right+1] = right - left
		left++
	}

	matchLen := 0
	A2 := A + A
	for i := 0; i < len(A2); i++ {
		for matchLen >= 0 && B[matchLen] != A2[i] {
			matchLen -= shifts[matchLen]
		}
		matchLen++
		if matchLen == N {
			return true
		}
	}

	return false
}
