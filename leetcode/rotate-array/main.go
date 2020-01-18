package main

func rotate(nums []int, k int) {
	l := len(nums)

	for i := 0; i < k; i++ {
		x := nums[l-1]
		for j := l - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = x
	}
}
