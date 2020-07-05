package main

import (
	"fmt"
	"math"
)

const maxArrayCount = 1000

func findTargetSumWays(nums []int, S int) int {
	memo := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		memo = append(memo, make([]int, maxArrayCount*2+1))
	}

	for i := 0; i < len(memo); i++ {
		for j := 0; j < len(memo[0]); j++ {
			memo[i][j] = math.MinInt32
		}
	}

	return calcurate(nums, 0, 0, S, memo)
}

func calcurate(nums []int, i, sum, S int, memo [][]int) int {
	if i == len(nums) {
		if sum == S {
			return 1
		} else {
			return 0
		}
	} else {
		if memo[i][sum+1000] != math.MinInt32 {
			return memo[i][sum+1000]
		}

		add := calcurate(nums, i+1, sum+nums[i], S, memo)
		subtract := calcurate(nums, i+1, sum-nums[i], S, memo)
		memo[i][sum+1000] = add + subtract
		return memo[i][sum+1000]
	}
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
