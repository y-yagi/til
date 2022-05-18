package main

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 0 {
		if n < 3 {
			return true
		}

		t := n % 2
		if t != 0 {
			return false
		}
		n /= 2
	}

	return true
}
