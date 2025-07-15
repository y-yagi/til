package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// tails[i] stores the smallest tail of all increasing subsequences of length i+1
	tails := make([]int, 0, len(nums))

	for _, num := range nums {
		// Binary search for the position to insert/replace
		left, right := 0, len(tails)
		for left < right {
			mid := (left + right) / 2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// If left == len(tails), append the number
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			// Replace the number at position left
			tails[left] = num
		}
	}

	return len(tails)
}
