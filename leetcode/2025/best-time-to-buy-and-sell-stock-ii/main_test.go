package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{in1: []int{7, 1, 5, 3, 6, 4}, want: 7},
	{in1: []int{1, 2, 3, 4, 5}, want: 4},
	{in1: []int{7, 6, 4, 3, 1}, want: 0},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := maxProfit(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
