package main

func fib(N int) int {
	if N == 1 || N == 0 {
		return N
	}

	return fib(N-1) + fib(N-2)
}
