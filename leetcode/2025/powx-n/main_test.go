package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  float64
	in2  int
	want float64
}{
	{2.00000, 10, 1024.00000},
	{2.10000, 3, 9.26100},
	{2.00000, -2, 0.25000},
	{0.00001, 2147483647, 0.00000},
	{1.00000, -2147483648, 1.00000},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		got := myPow(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
