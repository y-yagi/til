package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want bool
}{
	{[][]int{[]int{1, 5}, []int{8, 9}}, true},
	{[][]int{[]int{9, 19}, []int{4, 9}, []int{4, 17}}, false},
	{[][]int{[]int{5, 8}, []int{6, 8}}, false},
	{[][]int{[]int{13, 15}, []int{1, 13}}, true},
	{[][]int{[]int{0, 30}, []int{5, 10}, []int{15, 20}}, false},
	{[][]int{[]int{7, 10}, []int{2, 4}}, true},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := canAttendMeetings(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
