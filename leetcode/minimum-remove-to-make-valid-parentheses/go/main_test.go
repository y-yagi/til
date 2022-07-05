package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want string
}{
	{"lee(t(c)o)de)", "lee(t(c)o)de"},
	{"a)b(c)d", "ab(c)d"},
	{"))((", ""},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := minRemoveToMakeValid(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
