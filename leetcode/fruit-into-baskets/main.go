package main

import "fmt"

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}

func totalFruit(tree []int) int {
	answer := 0
	i := 0
	baskcets := map[int]int{}

	for j := 0; j < len(tree); j++ {
		baskcets[tree[j]]++

		for len(baskcets) >= 3 {
			baskcets[tree[i]]--
			if baskcets[tree[i]] == 0 {
				delete(baskcets, tree[i])
			}
			i++
		}
		answer = max(answer, j-i+1)
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
