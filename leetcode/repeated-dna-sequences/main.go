package main

func findRepeatedDnaSequences(s string) []string {
	dict := map[string]int{}
	ans := []string{}

	for i := 0; i < len(s)-9; i++ {
		sub := s[i : i+10]
		dict[sub]++
	}

	for k, v := range dict {
		if v >= 2 {
			ans = append(ans, k)
		}
	}

	return ans
}
