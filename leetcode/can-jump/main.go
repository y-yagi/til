package main

const (
	GOOD = iota
	BAD
	UNKNOWN
)

var memo []int

func canJump(nums []int) bool {
	memo = make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = UNKNOWN
	}
	memo[len(memo)-1] = GOOD
	return canJumpFromPosition(0, nums)
}

func canJumpFromPosition(pos int, nums []int) bool {
	if memo[pos] != UNKNOWN {
		if memo[pos] == GOOD {
			return true
		}
		return false
	}

	furthestJump := min(pos+nums[pos], len(nums)-1)

	for nextPos := furthestJump; nextPos > pos; nextPos-- {
		if canJumpFromPosition(nextPos, nums) {
			memo[pos] = GOOD
			return true
		}
	}

	memo[pos] = BAD
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
