package main

import "fmt"

func findMissingRanges(nums []int, lower int, upper int) []string {
	ans := []string{}
	next := lower

	for i := 0; i < len(nums); i++ {
		if nums[i] < next {
			continue
		}

		if nums[i] == next {
			next++
			continue
		}

		ans = append(ans, getRange(next, nums[i]-1))
		next = nums[i] + 1
	}

	if next <= upper {
		ans = append(ans, getRange(next, upper))
	}
	return ans
}

func getRange(n1, n2 int) string {
	if n1 == n2 {
		return fmt.Sprintf("%d", n1)
	}
	return fmt.Sprintf("%d->%d", n1, n2)
}
