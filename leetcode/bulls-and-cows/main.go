package main

import "fmt"

func getHint(secret string, guess string) string {
	dict := map[byte]int{}
	remain := []byte{}
	bulls := 0
	cows := 0

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			dict[secret[i]]++
			remain = append(remain, guess[i])
		}
	}

	for i := 0; i < len(remain); i++ {
		if v, found := dict[remain[i]]; found {
			if v > 0 {
				cows++
				dict[remain[i]]--
			}
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
