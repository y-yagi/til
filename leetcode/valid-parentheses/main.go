package main

func isValid(s string) bool {
	dict := map[string]string{")": "(", "}": "{", "]": "["}
	stack := make([]string, 0)

	for i := 0; i < len(s); i++ {
		index := string(s[i])
		_, found := dict[index]

		if found {
			top := ""
			if len(stack) == 0 {
				top = "#"
			} else {
				top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			}

			if top != dict[index] {
				return false
			}
		} else {
			stack = append(stack, index)
		}
	}

	return len(stack) == 0
}
