package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		dict[s[i]]++
		dict[t[i]]--
	}

	for _, v := range dict {
		if v != 0 {
			return false
		}
	}

	return true
}
