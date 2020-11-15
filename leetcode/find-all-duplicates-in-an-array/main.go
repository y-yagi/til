package main

func findDuplicates(nums []int) []int {
	ans := []int{}

	for _, n := range nums {
		if nums[abs(n)-1] < 0 {
			ans = append(ans, abs(n))
		}
		nums[abs(n)-1] *= -1
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}
