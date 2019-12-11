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
				top, stack = stack[0], stack[1:]
			}

			if top != dict[index] {
				return false
			}
		} else {
			stack = append([]string{index}, stack...)
		}
	}

	return len(stack) == 0
}
