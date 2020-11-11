package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	in2  []int
	in3  []string
	want []string
}{
	{[]string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"}, []string{"home", "about", "career"}},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := mostVisitedPattern(tt.in1, tt.in2, tt.in3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
