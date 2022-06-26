package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		sortColors(tt.in1)
		if !reflect.DeepEqual(tt.in1, tt.want) {
			t.Fatalf("got %v, want %v", tt.in1, tt.want)
		}
	}
}
