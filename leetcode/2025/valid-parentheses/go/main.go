package main

func isValid(s string) bool {
	dict := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, c := range s {
		if _, ok := dict[c]; ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || dict[stack[len(stack)-1]] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
