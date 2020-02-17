package main

func isSubsequence(s string, t string) bool {
	matched := 0
	j := 0
	for i := 0; i < len(s); i++ {
		for ; j < len(t); j++ {
			if s[i] == t[j] {
				matched++
				j++
				break
			}
		}
	}

	return matched == len(s)
}
