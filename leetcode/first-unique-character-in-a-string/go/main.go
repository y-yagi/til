package main

func firstUniqChar(s string) int {
	dict := map[byte]int{}

	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if v, found := dict[s[i]]; found && v == 1 {
			return i
		}
	}

	return -1
}
