package main

type HitCounter struct {
	times [300]int
	hits  [300]int
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (this *HitCounter) Hit(timestamp int) {
	index := timestamp % 300
	if this.times[index] != timestamp {
		this.times[index] = timestamp
		this.hits[index] = 1
	} else {
		this.hits[index]++
	}
}

func (this *HitCounter) GetHits(timestamp int) int {
	total := 0
	for i := 0; i < 300; i++ {
		if timestamp-this.times[i] < 300 {
			total += this.hits[i]
		}
	}
	return total
}
