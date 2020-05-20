package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want int
}{
	{[]string{"2", "1", "+", "3", "*"}, 9},
	{[]string{"4", "13", "5", "/", "+"}, 6},
	{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := evalRPN(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got '%v', want %v", tt.in1, got, tt.want)
		}
	}
}
