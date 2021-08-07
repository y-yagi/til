package main

import "testing"

func BenchmarkEncodeByGob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeByGob()
	}
}

func BenchmarkMarshalByJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarshalByJSON()
	}
}

func BenchmarkDecodeByGob(b *testing.B) {
	buf := EncodeByGob()
	for i := 0; i < b.N; i++ {
		DecodeByGob(buf)
	}
}

func BenchmarkUnmarshalByJSON(b *testing.B) {
	buf := MarshalByJSON()
	for i := 0; i < b.N; i++ {
		UnmarshalByJSON(buf)
	}
}
