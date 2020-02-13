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
	{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
	{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := intersection(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got, '%v', want '%v'", got, tt.want)
		}
	}
}
