package main

import (
	"sort"
)

type KthLargest struct {
	s []int
	k int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	l := len(nums)
	if k < l {
		l = k
	}
	obj := KthLargest{s: nums[:l], k: k}
	return obj
}

func (this *KthLargest) Add(val int) int {
	if len(this.s) == 0 {
		this.s = append(this.s, val)
		return this.s[this.k-1]
	}

	if len(this.s) >= this.k && this.s[this.k-1] >= val {
		return this.s[this.k-1]
	}

	if this.s[0] <= val {
		this.s = append([]int{val}, this.s[:this.k-1]...)
		return this.s[this.k-1]
	}

	i := len(this.s) / 2
	if this.s[i] > val {
		i = len(this.s) - 1
	}

	for ; i >= 0; i-- {
		if this.s[i] > val {
			this.s = append(this.s[:i+1], append([]int{val}, this.s[i+1:this.k-1]...)...)
			break
		}
	}
	return this.s[this.k-1]
}
