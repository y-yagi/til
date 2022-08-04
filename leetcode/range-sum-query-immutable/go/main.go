package main

type NumArray struct {
	nums []int
	sum  []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	n := &NumArray{nums: nums, sum: sum}
	return *n
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}
