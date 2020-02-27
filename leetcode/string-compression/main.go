package main

import "strconv"

func compress(chars []byte) int {
	anchor := 0
	write := 0

	for i := 0; i < len(chars); i++ {
		if i+1 == len(chars) || chars[i+1] != chars[i] {
			chars[write] = chars[anchor]
			write++
			if i > anchor {
				bs := []byte(strconv.Itoa(i - anchor + 1))
				for _, b := range bs {
					chars[write] = b
					write++
				}
			}
			anchor = i + 1
		}
	}

	return write
}
