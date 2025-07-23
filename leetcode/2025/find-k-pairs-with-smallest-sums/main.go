package main

import (
	"container/heap"
)

type pair struct {
	sum int
	i   int
	j   int
}

type minHeap []pair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return [][]int{}
	}

	h := &minHeap{}
	heap.Init(h)

	// Initialize heap with pairs from first element of nums1 with all elements of nums2
	for j := 0; j < len(nums2) && j < k; j++ {
		heap.Push(h, pair{sum: nums1[0] + nums2[j], i: 0, j: j})
	}

	result := make([][]int, 0, k)
	visited := make(map[[2]int]bool)

	for len(result) < k && h.Len() > 0 {
		current := heap.Pop(h).(pair)
		result = append(result, []int{nums1[current.i], nums2[current.j]})

		// Add next pair from the same row (i+1, j) if not visited
		if current.i+1 < len(nums1) {
			key := [2]int{current.i + 1, current.j}
			if !visited[key] {
				visited[key] = true
				heap.Push(h, pair{sum: nums1[current.i+1] + nums2[current.j], i: current.i + 1, j: current.j})
			}
		}
	}

	return result
}
