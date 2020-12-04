package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want [][]int
}{
	{
		[][]int{[]int{1, 91}, []int{1, 92}, []int{2, 93}, []int{2, 97}, []int{1, 60}, []int{2, 77}, []int{1, 65}, []int{1, 87}, []int{1, 100}, []int{2, 100}, []int{2, 76}},
		[][]int{[]int{1, 87}, []int{2, 88}},
	},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := highFive(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
