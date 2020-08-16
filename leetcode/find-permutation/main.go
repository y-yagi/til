package main

func findPermutation(s string) []int {
	ans := make([]int, len(s)+1)
	stack := []int{}
	j := 0

	for i := 1; i <= len(s); i++ {
		stack = append(stack, i)
		if s[i-1] == 'I' {
			for len(stack) != 0 {
				ans[j], stack = stack[len(stack)-1], stack[:len(stack)-1]
				j++
			}
		}
	}
	stack = append(stack, len(s)+1)
	for len(stack) != 0 {
		ans[j], stack = stack[len(stack)-1], stack[:len(stack)-1]
		j++
	}

	return ans
}
