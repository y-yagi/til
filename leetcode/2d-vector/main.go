package main

type Vector2D struct {
	v     [][]int
	inter int
	outer int
}

func Constructor(v [][]int) Vector2D {
	return Vector2D{v: v}
}

func (this *Vector2D) Next() int {
	if !this.HasNext() {
		return -1
	}

	v := this.v[this.inter][this.outer]
	this.outer++
	return v
}

func (this *Vector2D) HasNext() bool {
	this.move()
	return this.inter < len(this.v)
}

func (this *Vector2D) move() {
	if len(this.v) == 0 {
		return
	}

	for this.inter < len(this.v) && this.outer == len(this.v[this.inter]) {
		this.inter++
		this.outer = 0
	}
}
