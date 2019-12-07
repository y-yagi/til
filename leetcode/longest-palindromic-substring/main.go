package main

import "fmt"

func longestPalindrome(s string) string {
	var answer, tmp string

	for i := 0; i < len(s); i++ {
		index := string(s[i])
		if len(tmp) == 0 {
			tmp += index
			continue
		}

		if string(tmp[0]) == index {
			fmt.Printf("%s\n", tmp)
			tmp += index
			for j := 0; j < len(tmp); j++ {
				if string(tmp[j]) != string(s[i-j]) {
					break
				}
			}
			answer += tmp
		} else {
			tmp += index
		}
	}

	return answer
}

// func longestPalindrome(s string) string {
// 	if len(s) == 1 {
// 		return s
// 	} else if len(s) == 2 {
// 		if string(s[0]) == string(s[1]) {
// 			return string(s)
// 		} else {
// 			return string(s[0])
// 		}
// 	}
//
// 	palindromes := []string{}
// 	answer := ""
//
// 	for i := 0; i < len(s); i++ {
// 		for j, y := len(s[:i]), 0; j > 0; j, y = j-1, y+1 {
// 			if s[j] == s[y] {
// 				if j > y {
// 					palindromes = append(palindromes, s[y:j+1])
// 				} else if j == y {
// 					palindromes = append(palindromes, s[y:j+2])
// 				} else {
// 					palindromes = append(palindromes, s[j:y+1])
// 				}
// 			}
// 		}
// 	}
//
// 	for _, p := range palindromes {
// 		if len(p) > len(answer) {
// 			answer = p
// 		}
// 	}
//
// 	return answer
// }
