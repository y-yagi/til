package main

func isAlienSorted(words []string, order string) bool {
	dict := map[byte]int{}
	for i := 0; i < len(order); i++ {
		dict[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		before := words[i]
		after := words[i+1]
		l := min(len(before), len(after))
		j := 0

		for j = 0; j < l; j++ {
			if dict[before[j]] < dict[after[j]] {
				break
			} else if dict[before[j]] > dict[after[j]] {
				return false
			}
		}

		if j == l && len(before) > len(after) {
			return false
		}
	}

	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
