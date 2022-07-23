package main

import (
	"reflect"
	"testing"
)

//                          __________
// 				              ___ | Final step
//                 ___ | 20
//            ___ | 15
// _________ | 10
// First step

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{10, 15, 20}, 15},
	{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := minCostClimbingStairs(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
