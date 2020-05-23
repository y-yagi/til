package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{1, 0}, []int{0, 1}},
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := productExceptSelf(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got '%v', want %v", tt.in1, got, tt.want)
		}
	}
}
