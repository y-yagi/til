package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	results := [][]int{}
	counter := map[int]int{}
	for _, n := range nums {
		counter[n]++
	}

	comb := []int{}
	backtrace(comb, len(nums), counter, &results)
	return results
}

func backtrace(comb []int, N int, counter map[int]int, results *[][]int) {
	if len(comb) == N {
		*results = append(*results, comb)
		return
	}

	for num, count := range counter {
		if count == 0 {
			continue
		}

		comb = append([]int{num}, comb...)
		counter[num] = count - 1

		backtrace(comb, N, counter, results)

		comb = comb[1:]
		counter[num] = count
	}
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
