package main

func partitionLabels(S string) []int {
	dict := map[byte]int{}
	for i := 0; i < len(S); i++ {
		dict[S[i]] = i
	}

	var j, anchor int
	ans := []int{}
	for i := 0; i < len(S); i++ {
		j = max(j, dict[S[i]])
		if i == j {
			ans = append(ans, i-anchor+1)
			anchor = i + 1
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
