package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  string
	want string
}{
	{"I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"},
	{"The quick brown fox jumped over the lazy dog", "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := toGoatLatin(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
