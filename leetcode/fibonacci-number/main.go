package main

func fib(N int) int {
	if N <= 1 {
		return N
	}

	return memoize(N)
}

func memoize(N int) int {
	cache := map[int]int{0: 0, 1: 1}
	for i := 2; i <= N; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[N]
}
