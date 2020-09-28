package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	prod := 1
	var ans, left int
	for right, val := range nums {
		prod *= val
		for prod >= k {
			prod /= nums[left]
			left += 1
		}
		ans += right - left + 1
	}

	return ans
}
