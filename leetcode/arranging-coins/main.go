package main

func arrangeCoins(n int) int {
	s := 1
	a := 0

	for {
		n = n - s
		if n < 0 {
			break
		}
		a++
		s++
	}

	return a
}
