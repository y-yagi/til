package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)

	var tests = []struct {
		in   int
		want int
	}{
		{3, 4},
		{5, 5},
		{10, 5},
		{9, 8},
		{4, 8},
	}

	for _, tt := range tests {
		got := obj.Add(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestIsValid_WithEmptyStart(t *testing.T) {
	k := 1
	nums := []int{}
	obj := Constructor(k, nums)

	var tests = []struct {
		in   int
		want int
	}{
		{-3, -3},
		{-2, -2},
		{-4, -2},
		{0, 0},
		{4, 4},
	}

	for _, tt := range tests {
		got := obj.Add(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestIsValid_WithNegative(t *testing.T) {
	k := 2
	nums := []int{0}
	obj := Constructor(k, nums)

	var tests = []struct {
		in   int
		want int
	}{
		{-1, -1},
		{1, 0},
	}

	for _, tt := range tests {
		got := obj.Add(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestIsValid_WithOther(t *testing.T) {
	k := 3
	nums := []int{1, 1}
	obj := Constructor(k, nums)

	var tests = []struct {
		in   int
		want int
	}{
		{1, 1},
		{1, 1},
		{3, 1},
		{3, 1},
		{3, 3},
	}

	for _, tt := range tests {
		got := obj.Add(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
