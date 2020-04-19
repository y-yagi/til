package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	nums  []int
	snums []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	snums := make([]int, len(nums))
	copy(snums, nums)
	return Solution{nums: nums, snums: snums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.snums); i++ {
		this.swap(i, this.randRange(i, len(this.snums)))
	}
	return this.snums
}

func (this *Solution) swap(i, j int) {
	t := this.snums[i]
	this.snums[i] = this.snums[j]
	this.snums[j] = t
}

func (this *Solution) randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
