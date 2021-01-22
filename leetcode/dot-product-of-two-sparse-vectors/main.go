package main

type SparseVector struct {
	nums map[int]int
}

func Constructor(nums []int) SparseVector {
	s := SparseVector{nums: map[int]int{}}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			s.nums[i] = nums[i]
		}
	}

	return s
}

func (this *SparseVector) dotProduct(vec SparseVector) int {
	sum := 0
	for k, v := range this.nums {
		sum += vec.nums[k] * v
	}

	return sum
}
