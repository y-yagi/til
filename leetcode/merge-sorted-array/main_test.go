package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	in3  []int
	in4  int
	want []int
}{
	{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
}

func TestMere(t *testing.T) {
	for _, tt := range tests {
		merge(tt.in1, tt.in2, tt.in3, tt.in4)
		if !reflect.DeepEqual(tt.in1, tt.want) {
			t.Fatalf("got %v, want %v", tt.in1, tt.want)
		}
	}
}
