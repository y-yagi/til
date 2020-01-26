package main

type MyStack struct {
	values []int
}

func Constructor() MyStack {
	m := &MyStack{values: []int{}}
	return *m
}

func (this *MyStack) Push(x int) {
	this.values = append([]int{x}, this.values...)
}

func (this *MyStack) Pop() int {
	x := 0
	x, this.values = this.values[0], this.values[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.values[0]
}

func (this *MyStack) Empty() bool {
	return len(this.values) == 0
}
