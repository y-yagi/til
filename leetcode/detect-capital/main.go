package main

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	shouldAllUpper := false
	if isUpper(word[1]) {
		shouldAllUpper = true
		if !isUpper(word[0]) {
			return false
		}
	}

	for i := 2; i < len(word); i++ {
		if isUpper(word[i]) && !shouldAllUpper {
			return false
		}

		if !isUpper(word[i]) && shouldAllUpper {
			return false
		}
	}

	return true
}

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
