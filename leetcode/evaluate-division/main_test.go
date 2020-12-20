package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]string
	in2  []float64
	in3  [][]string
	want []float64
}{
	{[][]string{[]string{"a", "b"}, []string{"b", "c"}}, []float64{2.0, 3.0}, [][]string{[]string{"a", "c"}, []string{"b", "a"}, []string{"a", "e"}, []string{"a", "a"}, []string{"x", "x"}}, []float64{6.0, 0.5, -1.0, 1.0, -1.0}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := calcEquation(tt.in1, tt.in2, tt.in3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
