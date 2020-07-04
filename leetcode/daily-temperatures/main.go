package main

func dailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	stack := []int{}
	for i := len(T) - 1; i >= 0; i-- {
		for len(stack) != 0 && T[i] >= T[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}
