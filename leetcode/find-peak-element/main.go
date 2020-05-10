package main

func findPeakElement(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, l, r int) int {
	if l == r {
		return l
	}

	m := (l + r) / 2
	if nums[m] > nums[m+1] {
		return search(nums, l, m)
	}

	return search(nums, m+1, r)
}
