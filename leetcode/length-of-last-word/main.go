package main

func lengthOfLastWord(s string) int {
	answer, i := 0, len(s)-1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	for ; i >= 0 && s[i] != ' '; i-- {
		answer++
	}
	return answer
}
