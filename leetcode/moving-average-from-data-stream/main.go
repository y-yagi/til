package main

type MovingAverage struct {
	sum   int
	i     int
	queue []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{0, 0, make([]int, 0, size)}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) < cap(this.queue) {
		this.queue = append(this.queue, 0)
	}
	this.i = this.i % cap(this.queue)
	this.sum = this.sum - this.queue[this.i] + val
	this.queue[this.i] = val
	this.i++
	return float64(this.sum) / float64(len(this.queue))
}
