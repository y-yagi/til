package main

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
