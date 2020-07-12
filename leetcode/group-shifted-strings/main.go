package main

import "fmt"

func groupStrings(strings []string) [][]string {
	m := make(map[string][]string)
	result := make([][]string, 0)

	for i := 0; i < len(strings); i++ {
		cur := strings[i]
		diff := cur[0] - 'a'
		// Convert string to bytes and treat it as key.
		byteArr := make([]byte, 0)
		for k := 0; k < len(cur); k++ {
			if cur[k]-diff < 'a' {
				byteArr = append(byteArr, cur[k]-diff+26)
			} else {
				byteArr = append(byteArr, cur[k]-diff)
			}
		}
		m[string(byteArr)] = append(m[string(byteArr)], cur)
	}

	for _, val := range m {
		result = append(result, val)
	}
	return result
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
