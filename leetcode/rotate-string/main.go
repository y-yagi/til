package main

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if len(A) == 0 {
		return true
	}

	c := A[0]
	for i := 0; i < len(B); i++ {
		if B[i] == c {
			bi := i + 1
			ai := 1
			for ; ai < len(A); ai++ {
				if bi >= len(B) {
					bi = 0
				}
				if B[bi] != A[ai] {
					break
				}
				bi++
			}

			if ai == len(A) {
				return true
			}
		}
	}

	return false
}
