package main

import "sort"

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

type MinStack struct {
	values []int
	min    int
}

func Constructor() MinStack {
	values := []int{}
	m := &MinStack{values: values}
	m.min = maxInt
	return *m
}

func (m *MinStack) Push(x int) {
	m.values = append([]int{x}, m.values...)
	if x < m.min {
		m.min = x
	}
}

func (m *MinStack) Pop() {
	if len(m.values) == 0 {
		return
	}

	var x int
	x, m.values = m.values[0], m.values[1:]

	if len(m.values) == 0 {
		m.min = maxInt
	} else if x == m.min {
		a := make([]int, len(m.values))
		copy(a, m.values)
		sort.Ints(a)
		m.min = a[0]
	}
}

func (m *MinStack) Top() int {
	return m.values[0]
}

func (m *MinStack) GetMin() int {
	return m.min
}
