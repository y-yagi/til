package main

func isValid(s string) bool {
	dict := map[byte]byte{')': '(', '}': '{', ']': '['}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}

		pair, _ := dict[s[i]]
		if stack[len(stack)-1] == pair {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, s[i])
	}

	return len(stack) == 0
}
