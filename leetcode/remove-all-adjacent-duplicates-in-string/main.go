package main

func removeDuplicates(S string) string {
	stack := []byte{}
	for i := 0; i < len(S); i++ {
		if len(stack) != 0 && S[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}

	return string(stack)
}
