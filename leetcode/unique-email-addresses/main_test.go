package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	want int
}{
	{[]string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}, 2},
	{[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := numUniqueEmails(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
