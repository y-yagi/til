package main

import "fmt"

type data struct {
	sum   int
	count int
}

type start struct {
	station string
	t       int
}

type UndergroundSystem struct {
	in      map[int]start
	summary map[string]data
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{in: map[int]start{}, summary: map[string]data{}}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.in[id] = start{station: stationName, t: t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	s := this.in[id]
	k := fmt.Sprintf("%s-%s", s.station, stationName)
	if d, found := this.summary[k]; found {
		this.summary[k] = data{sum: d.sum + (t - s.t), count: d.count + 1}
	} else {
		this.summary[k] = data{sum: t - s.t, count: 1}
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	k := fmt.Sprintf("%s-%s", startStation, endStation)
	if d, found := this.summary[k]; found {
		return float64(d.sum) / float64(d.count)
	}

	return 0.0
}
