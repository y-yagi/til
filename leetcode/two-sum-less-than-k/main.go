package main

import "sort"

func twoSumLessThanK(A []int, K int) int {
	ans := -1
	sort.Ints(A)
	lo := 0
	hi := len(A) - 1

	for lo < hi {
		if A[lo]+A[hi] < K {
			ans = max(ans, A[lo]+A[hi])
			lo++
		} else {
			hi--
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
