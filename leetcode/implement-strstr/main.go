package main

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	if len(needle) == 0 {
		return 0
	}

	for i, j := 0, 0; i < len(haystack); i++ {
		if haystack[i] != needle[0] {
			continue
		}

		for j = 1; j < len(needle) && i+j < len(haystack); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}

		if j == len(needle) {
			return i
		}
	}

	return -1
}
