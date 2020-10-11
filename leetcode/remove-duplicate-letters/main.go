package main

func removeDuplicateLetters(s string) string {
	var stack []byte
	var counts [26]int
	var seen [26]bool

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']--
		if seen[s[i]-'a'] {
			continue
		}

		for j := len(stack) - 1; j >= 0; j-- {
			if !(s[i] < stack[j] && counts[stack[j]-'a'] > 0) {
				break
			}
			seen[stack[j]-'a'] = false
			stack = stack[:j]
		}
		seen[s[i]-'a'] = true
		stack = append(stack, s[i])
	}
	return string(stack)
}
