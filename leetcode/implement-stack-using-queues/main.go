package main

type MyStack struct {
	values []int
}

func Constructor() MyStack {
	m := &MyStack{}
	return *m
}

func (this *MyStack) Push(x int) {
	this.values = append(this.values, x)
}

func (this *MyStack) Pop() int {
	x := 0
	l := len(this.values) - 1
	x, this.values = this.values[l], this.values[:l]
	return x
}

func (this *MyStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.values) == 0
}
