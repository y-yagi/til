package main

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	return max(robber(nums, 0, n-2), robber(nums, 1, n-1))
}

func robber(nums []int, l, r int) int {
	var pre, cur int

	for i := l; i <= r; i++ {
		tmp := max(pre+nums[i], cur)
		pre = cur
		cur = tmp
	}

	return cur
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
