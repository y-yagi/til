package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []string
	in2  []string
	want []string
}{
	{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}, []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}},
	{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}, []string{"Shogun"}},
	{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}, []string{"Shogun"}},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := findRestaurant(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, in2 %v, got %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
