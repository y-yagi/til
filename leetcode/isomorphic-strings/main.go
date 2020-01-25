package main

func isIsomorphic(s string, t string) bool {
	dict1 := map[byte]byte{}
	dict2 := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		if v, ok := dict1[s[i]]; ok && t[i] != v {
			return false
		}

		if v, ok := dict2[t[i]]; ok && s[i] != v {
			return false
		}

		dict1[s[i]] = t[i]
		dict2[t[i]] = s[i]
	}

	return true
}
