package main

type MyQueue struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	q := &MyQueue{stack: []int{}}
	return *q
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	x := 0
	x, this.stack = this.stack[0], this.stack[1:]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
