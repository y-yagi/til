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
	{1, 2, 3},
	{-2, 3, 1},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := getSum(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got, '%v', want '%v'", got, tt.want)
		}
	}
}
