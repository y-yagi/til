package main

type MyCircularQueue struct {
	data   []int
	length int
	head   int
	tail   int
}

func Constructor(k int) MyCircularQueue {
	q := MyCircularQueue{
		data:   make([]int, k),
		length: 0,
		head:   0,
		tail:   -1,
	}
	return q
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	q.tail++
	if q.tail == cap(q.data) {
		q.tail = 0
	}

	q.data[q.tail] = value
	q.length++

	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.data[q.head] = 0

	q.head++
	if q.head == cap(q.data) {
		q.head = 0
	}

	q.length--

	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.head]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.tail]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.length == cap(q.data)
}
