package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	in2  int
	want int
}{
	{3, 2, 3},
	{7, 3, 28},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := uniquePaths(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
