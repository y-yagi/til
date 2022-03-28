package main

func moveZeroes(nums []int) {
	// All elements before the slow pointer (lastNonZeroFoundAt) are non-zeroes.
	// All elements between the current and slow pointer are zeroes.
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			t := nums[left]
			nums[left] = nums[i]
			nums[i] = t
			left++
		}
	}
}
