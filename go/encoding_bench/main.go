package main

import "fmt"

func main() {
	bufByGob := EncodeByGob()
	bytesByJSON := MarshalByJSON()
	fmt.Printf("EncodeByGob: %d\n", bufByGob.Len())
	fmt.Printf("MarshalByJSON: %d\n", len(bytesByJSON))
	// fmt.Printf("DecodeByGob: %+vn", DecodeByGob(bufByGob))
	// fmt.Printf("UnmarshalByJSON: %+v\n", UnmarshalByJSON(bytesByJSON))
}
