package main

func uniquePaths(m int, n int) int {
	cache := make([]int, m)
	cache[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j < len(cache); j++ {
			cache[j] += cache[j-1]
		}
	}
	return cache[len(cache)-1]
}
