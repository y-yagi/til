package main

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		celebrityCandidate := 0
		for i := 0; i < n; i++ {
			if knows(celebrityCandidate, i) {
				celebrityCandidate = i
			}
		}

		for j := 0; j < n; j++ {
			if celebrityCandidate == j {
				continue
			}

			if knows(celebrityCandidate, j) || !knows(j, celebrityCandidate) {
				return -1
			}
		}

		return celebrityCandidate
	}
}
