package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func plusOne(digits []int) []int {
	s := ""
	for _, b := range digits {
		s += strconv.Itoa(b)
	}

	i1, _ := new(big.Int).SetString(s, 0)
	i2, _ := new(big.Int).SetString("1", 0)
	sum := i1.Add(i1, i2)
	s = fmt.Sprintf("%v", sum)

	answer := []int{}
	for _, b := range s {
		i, _ := strconv.Atoi(string(b))
		answer = append(answer, i)
	}

	return answer
}
