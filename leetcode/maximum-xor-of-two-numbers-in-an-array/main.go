package main

import "strconv"

func findMaximumXOR(nums []int) int {
	maxNum := nums[0]
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}

	L := len(strconv.FormatInt(int64(maxNum), 2))
	var maxXor, currXor int
	prefixes := map[int]bool{}
	for i := L - 1; i > -1; i-- {
		// go to next next bit by the left shift
		maxXor <<= 1
		// set 1 in the smallest bit
		currXor = maxXor | 1
		prefixes = map[int]bool{}

		for _, num := range nums {
			prefixes[num>>i] = true
		}

		for p, _ := range prefixes {
			if _, found := prefixes[currXor^p]; found {
				maxXor = currXor
				break
			}
		}

	}

	return maxXor
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
