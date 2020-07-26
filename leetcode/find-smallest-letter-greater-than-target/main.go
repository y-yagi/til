package main

func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)
	for lo < hi {
		mi := lo + (hi-lo)/2
		if letters[mi] <= target {
			lo = mi + 1
		} else {
			hi = mi
		}
	}

	return letters[lo%len(letters)]
}
