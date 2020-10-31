package main

func findNumberOfLIS(nums []int) int {
	N := len(nums)
	if N <= 1 {
		return N
	}

	lengths := make([]int, N)
	counts := make([]int, N)
	for i := 0; i < len(counts); i++ {
		counts[i] = 1
	}

	for j := 0; j < N; j++ {
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				if lengths[i] >= lengths[j] {
					lengths[j] = lengths[i] + 1
					counts[j] = counts[i]
				} else if lengths[i]+1 == lengths[j] {
					counts[j] += counts[i]
				}
			}
		}
	}

	var longest, ans int
	for _, length := range lengths {
		longest = max(longest, length)
	}

	for i := 0; i < N; i++ {
		if lengths[i] == longest {
			ans += counts[i]
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
