package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	ans := 1

	for i := 1; i < len(dp); i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxval = max(maxval, dp[j])
			}
		}
		dp[i] = maxval + 1
		ans = max(ans, dp[i])
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
