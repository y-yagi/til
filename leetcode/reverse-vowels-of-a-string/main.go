package main

func reverseVowels(s string) string {
	vowels := map[byte]bool{'a': false, 'e': false, 'i': false, 'o': false, 'u': false, 'A': false, 'E': false, 'I': false, 'O': false, 'U': false}
	b := []byte(s)
	founds := []int{}

	for i := 0; i < len(b); i++ {
		if _, found := vowels[b[i]]; found {
			founds = append(founds, i)
		}
	}

	for i := 0; i < len(founds)/2; i++ {
		j := len(founds) - i - 1
		t := b[founds[i]]
		b[founds[i]] = b[founds[j]]
		b[founds[j]] = t
	}

	return string(b)
}
