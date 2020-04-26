package main

func findNumbers(nums []int) int {
	count := 0
	for _, n := range nums {
		if (n > 9 && n < 100) || (n > 999 && n < 10000) {
			count++
		}
	}
	return count
}
