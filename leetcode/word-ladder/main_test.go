package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	in3  []string
	want int
}{
	{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
	{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := ladderLength(tt.in1, tt.in2, tt.in3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("got %v, want %v", got, tt.want)
		}
	}
}
