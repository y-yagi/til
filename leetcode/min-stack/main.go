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
	m.stack = append(m.stack, x)
	if len(m.minStack) == 0 || x <= m.minStack[len(m.minStack)-1] {
		m.minStack = append(m.minStack, x)
	}
}

func (m *MinStack) Pop() {
	var x int
	l := len(m.stack)
	minl := len(m.minStack)
	x, m.stack = m.stack[l-1], m.stack[:l-1]
	if x == m.minStack[minl-1] {
		m.minStack = m.minStack[:minl-1]
	}
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}
