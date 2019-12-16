package main

import "strconv"

func countAndSay(n int) string {
	result := "1"

	for i := 1; i < n; i++ {
		newResult := ""
		j := 0

		for j < len(result) {
			currentCharCount := 1
			currentChar := result[j]

			for j < len(result)-1 && result[j] == result[j+1] {
				currentCharCount++
				j++
			}
			j++
			newResult += strconv.Itoa(currentCharCount) + string([]byte{currentChar})
		}

		result = newResult
	}

	return result
}
