package main

type MaxStack struct {
	stack  []int
	mStack []int
}

func Constructor() MaxStack {
	m := &MaxStack{}
	return *m
}

func (m *MaxStack) Push(x int) {
	var max int
	if len(m.mStack) == 0 {
		max = x
	} else {
		max = m.mStack[len(m.mStack)-1]
	}

	if x > max {
		max = x
	}

	m.mStack = append(m.mStack, max)
	m.stack = append(m.stack, x)
}

func (m *MaxStack) Pop() int {
	var x int
	l := len(m.stack)
	m.mStack = m.mStack[:len(m.mStack)-1]
	x, m.stack = m.stack[l-1], m.stack[:l-1]
	return x
}

func (m *MaxStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MaxStack) PeekMax() int {
	return m.mStack[len(m.mStack)-1]
}

func (m *MaxStack) PopMax() int {
	max := m.PeekMax()
	buffer := []int{}
	for m.Top() != max {
		buffer = append(buffer, m.Pop())
	}
	m.Pop()
	for i := 0; i < len(buffer); i++ {
		m.Push(buffer[i])
	}
	return max
}
