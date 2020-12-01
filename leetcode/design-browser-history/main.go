package main

type BrowserHistory struct {
	queue   []string
	current int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{queue: []string{homepage}}
}

func (this *BrowserHistory) Visit(url string) {
	this.queue = this.queue[:this.current+1]
	this.queue = append(this.queue, url)
	this.current = len(this.queue) - 1
}

func (this *BrowserHistory) Back(steps int) string {
	this.current = this.current - steps
	if this.current < 0 {
		this.current = 0
	}

	return this.queue[this.current]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.current = this.current + steps
	if this.current > len(this.queue)-1 {
		this.current = len(this.queue) - 1
	}

	return this.queue[this.current]
}
