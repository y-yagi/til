package main

func firstUniqChar(s string) int {
	dict := make(map[rune]int)

	for _, c := range s {
		dict[c]++
	}

	for i, c := range s {
		if v, found := dict[c]; found && v == 1 {
			return i
		}
	}

	return -1
}
