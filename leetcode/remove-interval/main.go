package main

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	ans := [][]int{}

	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		if end <= toBeRemoved[0] || start >= toBeRemoved[1] {
			ans = append(ans, interval)
		} else if start < toBeRemoved[0] && end > toBeRemoved[1] {
			ans = append(ans, []int{start, toBeRemoved[0]})
			ans = append(ans, []int{toBeRemoved[1], end})
		} else if start < toBeRemoved[0] && end <= toBeRemoved[1] {
			ans = append(ans, []int{start, toBeRemoved[0]})
		} else if start >= toBeRemoved[0] && end > toBeRemoved[1] {
			ans = append(ans, []int{toBeRemoved[1], end})
		}
	}

	return ans
}
