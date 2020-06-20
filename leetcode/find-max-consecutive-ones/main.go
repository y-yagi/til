package main

func findMaxConsecutiveOnes(nums []int) int {
	var ans, left, right int

	for i := 0; i < len(nums); i++ {
		right++
		if nums[i] == 0 {
			left = right
			right = 0
		}
		ans = max(ans, left+right)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
