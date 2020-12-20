package main

func maxNumberOfBalloons(text string) int {
	dict := map[byte]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}

	for i := 0; i < len(text); i++ {
		if _, found := dict[text[i]]; found {
			dict[text[i]]++
		}
	}
	ans := dict['b']
	ans = min(ans, dict['a'])
	ans = min(ans, dict['n'])
	ans = min(ans, dict['l']/2)
	ans = min(ans, dict['o']/2)

	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
