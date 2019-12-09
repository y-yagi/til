package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	answer := ""
	min := len(strs[0])

	for _, str := range strs[0:] {
		if len(str) < min {
			min = len(str)
		}
	}

out:
	for i := 0; i < min; i++ {
		index := string(strs[0][i])

		for _, str := range strs[0:] {
			if index != string(str[i]) {
				break out
			}
		}
		answer += index
	}

	return answer
}
