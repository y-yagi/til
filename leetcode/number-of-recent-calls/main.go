package main

type RecentCounter struct {
	queue []int
	count int
}

func Constructor() RecentCounter {
	return RecentCounter{queue: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	this.count++

	if len(this.queue) < 2 {
		return this.count
	}

	min := this.queue[0]
	max := this.queue[this.count-1]
	for max-min > 3000 {
		this.queue = this.queue[1:]
		this.count--
		min = this.queue[0]
	}

	return this.count
}
