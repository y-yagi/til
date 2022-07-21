package main

import "strings"

func decodeString(s string) string {
	stackNums := make([]int, 0)
	stackStr := make([]string, 0)

	var ans string
	var num int

	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			num = 10*num + int(s[i]) - '0'
		case s[i] == '[':
			stackNums = append(stackNums, num)
			num = 0
			stackStr = append(stackStr, ans)
			ans = ""
		case s[i] == ']':
			tmp := stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]
			count := stackNums[len(stackNums)-1]
			stackNums = stackNums[:len(stackNums)-1]
			ans = tmp + strings.Repeat(ans, count)
		default:
			ans += string(s[i])
		}
	}
	return ans
}
