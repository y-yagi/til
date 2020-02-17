package main

func findTheDifference(s string, t string) byte {
	dict := make([]int, 26)

	for i := 0; i < len(s); i++ {
		dict[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		if dict[t[i]-'a'] == 0 {
			return t[i]
		} else {
			dict[t[i]-'a']--
		}
	}

	return 'A'
}
