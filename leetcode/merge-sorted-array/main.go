package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[0:m]
	nums1 = append(nums1, nums2[0:n]...)
	sort.Ints(nums1)
}
