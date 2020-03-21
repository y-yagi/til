package main

func maxSubArray(nums []int) int {
	ans := nums[0]
	cur := nums[0]
	for p := 1; p < len(nums); p++ {
		if cur >= 0 {
			cur += nums[p]
		} else {
			cur = nums[p]
		}

		ans = max(ans, cur)
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
