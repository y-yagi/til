package main

func maxDistToClosest(seats []int) int {
	N := len(seats)
	K := 0
	ans := 0

	for i := 0; i < N; i++ {
		if seats[i] == 1 {
			K = 0
		} else {
			K++
			ans = max(ans, (K+1)/2)
		}
	}

	for i := 0; i < N; i++ {
		if seats[i] == 1 {
			ans = max(ans, i)
			break
		}
	}

	for i := N - 1; i >= 0; i-- {
		if seats[i] == 1 {
			ans = max(ans, N-1-i)
			break
		}
	}

	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
