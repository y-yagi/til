package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []string
}{
	{[]int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
	{[]int{10, 3, 8, 9, 4}, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findRelativeRanks(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
