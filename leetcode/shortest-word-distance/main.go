package main

func shortestDistance(words []string, word1 string, word2 string) int {
	i1 := -1
	i2 := -1
	minDistance := len(words)

	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			i1 = i
		}
		if words[i] == word2 {
			i2 = i
		}

		if i1 != -1 && i2 != -1 {
			minDistance = min(minDistance, abs(i1-i2))
		}
	}

	return minDistance
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
