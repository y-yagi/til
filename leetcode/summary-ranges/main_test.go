package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []string
}{
	{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
	{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := summaryRanges(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
