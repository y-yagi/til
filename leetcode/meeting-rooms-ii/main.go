package main

import (
	"container/heap"
	"sort"
)

type Allocator []int

func (h Allocator) Len() int           { return len(h) }
func (h Allocator) Less(i, j int) bool { return h[i] < h[j] }
func (h Allocator) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Allocator) Peek() int          { return h[0] }

func (h *Allocator) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Allocator) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	allocator := &Allocator{}
	heap.Init(allocator)
	heap.Push(allocator, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= allocator.Peek() {
			heap.Pop(allocator)
		}
		heap.Push(allocator, intervals[i][1])
	}

	return allocator.Len()
}
