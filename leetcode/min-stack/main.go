package main

import (
	"math"
)

const maxInt = math.MaxInt64

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	m := &MinStack{}
	return *m
}

func (m *MinStack) Push(x int) {
	m.stack = append([]int{x}, m.stack...)
	if len(m.minStack) == 0 || x <= m.minStack[0] {
		m.minStack = append([]int{x}, m.minStack...)
	}
}

func (m *MinStack) Pop() {
	var x int
	x, m.stack = m.stack[0], m.stack[1:]
	if x == m.minStack[0] {
		m.minStack = m.minStack[1:]
	}
}

func (m *MinStack) Top() int {
	return m.stack[0]
}

func (m *MinStack) GetMin() int {
	return m.minStack[0]
}
