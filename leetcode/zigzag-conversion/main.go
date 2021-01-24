package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	ret := []byte{}
	n := len(s)
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			ret = append(ret, s[i+j])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				ret = append(ret, s[j+cycleLen-i])
			}
		}
	}

	return string(ret)
}
