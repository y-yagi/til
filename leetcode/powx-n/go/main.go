package main

func myPow(x float64, n int) float64 {
	var temp float64
	if n == 0 {
		return 1
	}

	temp = myPow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}
	if n > 0 {
		return x * temp * temp
	}
	return (temp * temp) / x
}
