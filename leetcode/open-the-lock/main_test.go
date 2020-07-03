package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	in2  string
	want int
}{
	{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
	{[]string{"8888"}, "0009", 1},
	{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
	{[]string{"0000"}, "8888", -1},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := openLock(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
