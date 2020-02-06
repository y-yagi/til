package main

func moveZeroes(nums []int) {
	l := len(nums)
	zeros := 0
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			zeros++
			i--
			l--
		}
	}

	for i := 0; i < zeros; i++ {
		nums = append(nums, 0)
	}
}
