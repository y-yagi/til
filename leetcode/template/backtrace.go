package main

func bactrace(candidate int) {
	if findSolution(candidate) {
		output(candidate)
	}

	for _, next_candidate := range list_of_candidate {
		if isValid(next_candidate) {
			// try this partial candidate solution
			place(next_candidate)
			// given the candidate, explore further.
			backtrack(next_candidate)
			// backtrack
			remove(next_candidate)
		}
	}
}

// Real example
func combine(n int, k int) [][]int {
	ans := [][]int{}
	curr := []int{}
	backtrace(n, k, 1, curr, &ans)
	return ans
}

func backtrace(n, k, first int, curr []int, ans *[][]int) {
	if len(curr) == k {
		b := make([]int, len(curr))
		copy(b, curr)
		*ans = append(*ans, b)
	}

	for i := first; i < n+1; i++ {
		curr = append(curr, i)
		backtrace(n, k, i+1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}
