package main

import (
	"math/rand"
)

type Solution struct {
	prefixSums []int
	totalSum   int
}

func Constructor(w []int) Solution {
	s := Solution{}
	s.prefixSums = make([]int, len(w))

	prefixSum := 0
	for i := 0; i < len(w); i++ {
		prefixSum += w[i]
		s.prefixSums[i] = prefixSum
	}
	s.totalSum = prefixSum

	return s
}

func (this *Solution) PickIndex() int {
	target := float64(this.totalSum) * rand.Float64()
	i := 0
	for ; i < len(this.prefixSums); i++ {
		if int(target) < this.prefixSums[i] {
			return i
		}
	}
	return i - 1
}
