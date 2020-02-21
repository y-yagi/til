package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want []string
}{
	{1, []string{"1"}},
	{3, []string{"1", "2", "Fizz"}},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := fizzBuzz(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
