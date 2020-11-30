package main

func countArrangement(N int) int {
	visited := make([]bool, N+1)
	ans := 0
	calcurate(N, 1, visited, &ans)
	return ans
}

func calcurate(N, pos int, visited []bool, ans *int) {
	if pos > N {
		*ans++
	}

	for i := 1; i <= N; i++ {
		if !visited[i] && (pos%i == 0 || i%pos == 0) {
			visited[i] = true
			calcurate(N, pos+1, visited, ans)
			visited[i] = false
		}
	}
}
