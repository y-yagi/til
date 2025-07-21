package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  []int
	want []int
}{
	{in1: []int{1, 2, 2, 1}, in2: []int{2, 2}, want: []int{2}},
	{in1: []int{4, 9, 5}, in2: []int{9, 4, 9, 8, 4}, want: []int{9, 4}},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := intersection(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
