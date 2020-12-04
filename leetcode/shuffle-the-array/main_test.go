package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want []int
}{
	{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
	{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := shuffle(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
