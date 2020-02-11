package main

func reverseString(s []byte) {
	m := len(s) / 2
	l := len(s) - 1

	for i := 0; i < m; i++ {
		t := s[l-i]
		s[l-i] = s[i]
		s[i] = t
	}
}
