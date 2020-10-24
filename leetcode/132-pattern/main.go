package main

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	stack := []int{}
	mins := make([]int, len(nums))
	mins[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		mins[i] = min(mins[i-1], nums[i])
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] > mins[j] {
			for len(stack) > 0 && stack[len(stack)-1] <= mins[j] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 && stack[len(stack)-1] < nums[j] {
				return true
			}
			stack = append(stack, nums[j])
		}
	}

	return false
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
