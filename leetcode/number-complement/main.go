package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	bits := fmt.Sprintf("%b", num)
	reversed := ""
	for _, b := range bits {
		if b == '1' {
			reversed += "0"
		} else {
			reversed += "1"
		}
	}

	v, _ := strconv.ParseInt(reversed, 2, 64)
	return int(v)
}
