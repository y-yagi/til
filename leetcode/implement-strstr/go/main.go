package main

func strStr(haystack string, needle string) int {
	n := len(haystack)
	l := len(needle)

	for i := 0; i < n-l+1; i++ {
		if haystack[i:i+l] == needle {
			return i
		}
	}

	return -1
}
