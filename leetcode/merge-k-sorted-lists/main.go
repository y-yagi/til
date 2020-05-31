package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h *NodeHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
	mergeHeap := &NodeHeap{}
	heap.Init(mergeHeap)

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(mergeHeap, lists[i])
		}
	}
	merged := &ListNode{0, nil}
	head := merged
	for mergeHeap.Len() > 0 {
		next := heap.Pop(mergeHeap).(*ListNode)
		if next.Next != nil {
			heap.Push(mergeHeap, next.Next)
		}
		head.Next = next
		head = head.Next
	}
	return merged.Next
}
