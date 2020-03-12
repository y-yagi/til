package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	answer := ""
	l := len(s)
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < l; j += cycleLen {
			answer += string(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < l {
				answer += string(s[j+cycleLen-i])
			}
		}
	}
	return answer
}
