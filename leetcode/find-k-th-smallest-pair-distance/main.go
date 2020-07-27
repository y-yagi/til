package main

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	lo := 0
	hi := nums[len(nums)-1] - nums[0]
	for lo < hi {
		mi := (lo + hi) / 2
		count, left := 0, 0
		for right := 0; right < len(nums); right++ {
			for nums[right]-nums[left] > mi {
				left++
			}
			count += right - left
		}
		if count >= k {
			hi = mi
		} else {
			lo = mi + 1
		}
	}

	return lo
}
