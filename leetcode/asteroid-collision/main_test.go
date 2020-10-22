package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{5, 10, -5}, []int{5, 10}},
	{[]int{8, -8}, []int{}},
	{[]int{10, 2, -5}, []int{10}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := asteroidCollision(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
