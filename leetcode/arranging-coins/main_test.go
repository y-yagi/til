package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want int
}{
	{5, 2},
	{8, 3},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := arrangeCoins(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
