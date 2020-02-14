package main

func isPerfectSquare(num int) bool {
	if num < 1 {
		return false
	}

	if num == 1 {
		return true
	}

	t := num / 2
	for t*t > num {
		t = (t + num/t) / 2
	}
	return t*t == num
}
