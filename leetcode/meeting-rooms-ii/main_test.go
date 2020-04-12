package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{[][]int{[]int{9, 10}, []int{4, 9}, []int{4, 17}}, 2},
	{[][]int{}, 0},
	{[][]int{[]int{0, 30}, []int{5, 10}, []int{15, 20}}, 2},
	{[][]int{[]int{7, 10}, []int{2, 4}}, 1},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := minMeetingRooms(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
